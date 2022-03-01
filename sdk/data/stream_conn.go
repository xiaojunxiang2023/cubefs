// Copyright 2018 The Chubao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package data

import (
	"fmt"
	"hash/crc32"
	"net"
	"sort"
	"time"

	"github.com/chubaofs/chubaofs/proto"
	"github.com/chubaofs/chubaofs/util"
	"github.com/chubaofs/chubaofs/util/errors"
	"github.com/chubaofs/chubaofs/util/log"
	"github.com/chubaofs/chubaofs/util/tracing"
)

const (
	StreamSendReadMaxRetry      = 3
	StreamSendOverWriteMaxRetry = 1
	StreamSendOverWriteTimeout  = 2 * time.Second

	StreamSendSleepInterval = 100 * time.Millisecond
	StreamSendTimeout       = 2 * time.Minute

	StreamReadConsistenceRetry   = 50
	StreamReadConsistenceTimeout = 1 * time.Minute

	IdleConnTimeoutData  = 30
	ConnectTimeoutDataMs = 500
	ReadTimeoutData      = 3
	WriteTimeoutData     = 3

	hostErrAccessTimeout = 300 // second
)

type GetReplyFunc func(conn *net.TCPConn) (err error, again bool)

// StreamConn defines the struct of the stream connection.
type StreamConn struct {
	dp       *DataPartition
	currAddr string
}

var StreamConnPool *util.ConnectPool

// NewStreamConn returns a new stream connection.
func NewStreamConn(dp *DataPartition, follower bool) *StreamConn {
	if !follower {
		return &StreamConn{
			dp:       dp,
			currAddr: dp.LeaderAddr,
		}
	}

	if dp.ClientWrapper.CrossRegionHATypeQuorum() {
		return &StreamConn{
			dp:       dp,
			currAddr: dp.getNearestCrossRegionHost(),
		}
	}

	if dp.ClientWrapper.NearRead() {
		return &StreamConn{
			dp:       dp,
			currAddr: dp.getNearestHost(),
		}
	}

	return &StreamConn{
		dp:       dp,
		currAddr: dp.getFollowerReadHost(),
	}
}

func NewStreamConnWithAddr(dp *DataPartition, addr string) *StreamConn {
	return &StreamConn{
		dp:       dp,
		currAddr: addr,
	}
}

// String returns the string format of the stream connection.
func (sc *StreamConn) String() string {
	if sc == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Partition(%v) CurrentAddr(%v) Hosts(%v)", sc.dp.PartitionID, sc.currAddr, sc.dp.Hosts)
}

func (sc *StreamConn) sendToDataPartition(req *Packet) (conn *net.TCPConn, err error) {
	var tracer = tracing.TracerFromContext(req.Ctx()).ChildTracer("StreamConn.sendToDataPartition").
		SetTag("req.ReqID", req.ReqID).
		SetTag("req.Op", req.GetOpMsg()).
		SetTag("req.PartitionID", req.PartitionID).
		SetTag("req.ExtentID", req.ExtentID).
		SetTag("req.Size", req.Size).
		SetTag("req.ExtentOffset", req.ExtentOffset)
	defer func() {
		if conn != nil {
			tracer.SetTag("ret.conn.local", conn.LocalAddr().String())
			tracer.SetTag("ret.conn.remote", conn.RemoteAddr().String())
		}
		tracer.SetTag("ret.err", err)
		tracer.Finish()
	}()
	req.SetCtx(tracer.Context())

	log.LogDebugf("sendToDataPartition: send to addr(%v), reqPacket(%v)", sc.currAddr, req)
	if conn, err = StreamConnPool.GetConnect(sc.currAddr); err != nil {
		log.LogWarnf("sendToDataPartition: get connection to curr addr failed, addr(%v) reqPacket(%v) err(%v)", sc.currAddr, req, err)
		return
	}
	if err = func() error {
		var tracer = tracing.TracerFromContext(req.Ctx()).ChildTracer("dataPartition.sendToDataPartition[WriteToConn]").
			SetTag("reqID", req.GetReqID()).
			SetTag("reqOp", req.GetOpMsg())
		defer tracer.Finish()
		return req.WriteToConnNs(conn, sc.dp.ClientWrapper.connConfig.WriteTimeoutNs)
	}(); err != nil {
		log.LogWarnf("sendToDataPartition: failed to write to addr(%v) err(%v)", sc.currAddr, err)
		return
	}
	log.LogDebugf("sendToDataPartition exit: send to addr(%v) reqPacket(%v) successfully", sc.currAddr, req)
	return
}

// sortByStatus will return hosts list sort by host status for DataPartition.
// The order from front to back is "status(true)/status(false)/failedHost".
func sortByStatus(dp *DataPartition, failedHost string) (hosts []string) {
	var inactiveHosts []string
	hostsStatus := dp.ClientWrapper.HostsStatus
	var dpHosts []string
	if dp.ClientWrapper.CrossRegionHATypeQuorum() {
		dpHosts = dp.getSortedCrossRegionHosts()
	} else if dp.ClientWrapper.FollowerRead() && dp.ClientWrapper.NearRead() {
		dpHosts = dp.NearHosts
	}
	if len(dpHosts) == 0 {
		dpHosts = dp.Hosts
	}

	for _, addr := range dpHosts {
		if addr == failedHost {
			continue
		}
		status, ok := hostsStatus[addr]
		if ok {
			if status {
				hosts = append(hosts, addr)
			} else {
				inactiveHosts = append(inactiveHosts, addr)
			}
		} else {
			inactiveHosts = append(inactiveHosts, addr)
			log.LogWarnf("sortByStatus: can not find host[%v] in HostsStatus, dp[%d]", addr, dp.PartitionID)
		}
	}

	sortByAccessErrTs(dp, hosts)

	hosts = append(hosts, inactiveHosts...)
	hosts = append(hosts, failedHost)

	log.LogDebugf("sortByStatus: dp(%v) sortedHost(%v) failedHost(%v)", dp, hosts, failedHost)

	return
}

func sortByAccessErrTs(dp *DataPartition, hosts []string) {

	for _, host := range hosts {
		ts, ok := dp.hostErrMap.Load(host)
		if ok && time.Now().UnixNano()-ts.(int64) > hostErrAccessTimeout*1e9 {
			dp.hostErrMap.Delete(host)
		}
	}

	sort.Slice(hosts, func(i, j int) bool {
		var iTime, jTime int64
		iTs, ok := dp.hostErrMap.Load(hosts[i])
		if ok {
			iTime = iTs.(int64)
		}
		jTs, ok := dp.hostErrMap.Load(hosts[j])
		if ok {
			jTime = jTs.(int64)
		}
		return iTime < jTime
	})
}

func (dp *DataPartition) getNearestHost() string {
	hostsStatus := dp.ClientWrapper.HostsStatus
	for _, addr := range dp.NearHosts {
		status, ok := hostsStatus[addr]
		if ok {
			if !status {
				continue
			}
		}
		return addr
	}
	return dp.LeaderAddr
}

func (dp *DataPartition) getFollowerReadHost() string {
	if len(dp.Hosts) > 0 {
		err, host := dp.getEpochReadHost(dp.Hosts)
		if err == nil {
			return host
		}
	}
	return dp.LeaderAddr
}

func (sc *StreamConn) getReadReply(conn *net.TCPConn, reqPacket *Packet, req *ExtentRequest) (readBytes int, reply *Packet, tryOther bool, err error) {
	var tracer = tracing.TracerFromContext(reqPacket.Ctx()).ChildTracer("StreamConn.getReadReply").
		SetTag("remote", conn.RemoteAddr().String())
	defer tracer.Finish()

	readBytes = 0
	for readBytes < int(reqPacket.Size) {
		replyPacket := NewReply(reqPacket.Ctx(), reqPacket.ReqID, reqPacket.PartitionID, reqPacket.ExtentID)
		bufSize := util.Min(util.ReadBlockSize, int(reqPacket.Size)-readBytes)
		replyPacket.Data = req.Data[readBytes : readBytes+bufSize]
		e := replyPacket.readFromConn(conn, sc.dp.ClientWrapper.connConfig.ReadTimeoutNs)
		if e != nil {
			log.LogWarnf("getReadReply: failed to read from connect, ino(%v) req(%v) readBytes(%v) err(%v)", reqPacket.inode, reqPacket, readBytes, e)
			// Upon receiving TryOtherAddrError, other hosts will be retried.
			return readBytes, replyPacket, true, e
		}
		//log.LogDebugf("ExtentReader Read: ResultCode(%v) req(%v) reply(%v) readBytes(%v)", replyPacket.GetResultMsg(), reqPacket, replyPacket, readBytes)

		e = checkReadReplyValid(reqPacket, replyPacket)
		if e != nil {
			// Dont change the error message, since the caller will
			// check if it is NotLeaderErr.
			return readBytes, replyPacket, false, e
		}

		readBytes += int(replyPacket.Size)
	}
	return readBytes, nil, false, nil
}

func checkReadReplyValid(request *Packet, reply *Packet) (err error) {
	if reply.ResultCode != proto.OpOk {
		err = errors.New(fmt.Sprintf("checkReadReplyValid: ResultCode(%v) NOK", reply.GetResultMsg()))
		return
	}
	if !request.isValidReadReply(reply) {
		err = errors.New(fmt.Sprintf("checkReadReplyValid: inconsistent req and reply, req(%v) reply(%v)", request, reply))
		return
	}
	expectCrc := crc32.ChecksumIEEE(reply.Data[:reply.Size])
	if reply.CRC != expectCrc {
		err = errors.New(fmt.Sprintf("checkReadReplyValid: inconsistent CRC, expectCRC(%v) replyCRC(%v)", expectCrc, reply.CRC))
		return
	}
	return nil
}
