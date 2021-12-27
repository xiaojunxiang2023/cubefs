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

package master

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/chubaofs/chubaofs/raftstore"
	"github.com/tiglabs/raft/proto"
)

//config key
const (
	colonSplit = ":"
	commaSplit = ","
	cfgPeers   = "peers"
	// if the data partition has not been reported within this interval  (in terms of seconds), it will be considered as missing.
	missingDataPartitionInterval        = "missingDataPartitionInterval"
	dataPartitionTimeOutSec             = "dataPartitionTimeOutSec"
	NumberOfDataPartitionsToLoad        = "numberOfDataPartitionsToLoad"
	secondsToFreeDataPartitionAfterLoad = "secondsToFreeDataPartitionAfterLoad"
	nodeSetCapacity                     = "nodeSetCap"
	cfgMetaNodeReservedMem              = "metaNodeReservedMem"
	heartbeatPortKey                    = "heartbeatPort"
	replicaPortKey                      = "replicaPort"
)

//default value
const (
	defaultTobeFreedDataPartitionCount         = 1000
	defaultSecondsToFreeDataPartitionAfterLoad = 5 * 60 // a data partition can only be freed after loading 5 mins
	defaultIntervalToFreeDataPartition         = 10     // in terms of seconds
	defaultIntervalToCheckHeartbeat            = 60
	defaultIntervalToCheckDataPartition        = 60
	defaultIntervalToCheckCrc                  = 20 * defaultIntervalToCheckHeartbeat // in terms of seconds
	noHeartBeatTimes                           = 3                                    // number of times that no heartbeat reported
	defaultNodeTimeOutSec                      = noHeartBeatTimes * defaultIntervalToCheckHeartbeat
	defaultDataPartitionTimeOutSec             = 10 * defaultIntervalToCheckHeartbeat
	defaultMissingDataPartitionInterval        = 24 * 3600

	defaultIntervalToAlarmMissingDataPartition = 60 * 60
	timeToWaitForResponse                      = 120         // time to wait for response by the master during loading partition
	defaultPeriodToLoadAllDataPartitions       = 60 * 60 * 4 // how long we need to load all the data partitions on the master every time
	defaultNumberOfDataPartitionsToLoad        = 50          // how many data partitions to load every time
	defaultMetaPartitionTimeOutSec             = 10 * defaultIntervalToCheckHeartbeat
	//DefaultMetaPartitionMissSec                         = 3600

	defaultIntervalToAlarmMissingMetaPartition         = 10 * 60 // interval of checking if a replica is missing
	defaultMetaPartitionMemUsageThreshold      float32 = 0.75    // memory usage threshold on a meta partition
	defaultMaxMetaPartitionCountOnEachNode             = 10000
	defaultReplicaNum                                  = 3
	defaultDiffSpaceUsage                              = 10 * 1024 * 1024 * 1024
	defaultCrossZoneNum                                = 3
)

// AddrDatabase is a map that stores the address of a given host (e.g., the leader)
var AddrDatabase = make(map[uint64]string)

type clusterConfig struct {
	secondsToFreeDataPartitionAfterLoad int64
	NodeTimeOutSec                      int64
	MissingDataPartitionInterval        int64
	DataPartitionTimeOutSec             int64
	IntervalToAlarmMissingDataPartition int64
	PeriodToLoadALLDataPartitions       int64
	metaNodeReservedMem                 uint64
	IntervalToCheckDataPartition        int // seconds
	numberOfDataPartitionsToFree        int
	numberOfDataPartitionsToLoad        int
	nodeSetCapacity                     int
	MetaNodeThreshold                   float32
	MetaNodeDeleteBatchCount            uint64 //metanode delete batch count
	MetaNodeReqRateLimit                uint64
	MetaNodeReqOpRateLimitMap           map[uint8]uint64
	DataNodeReqZoneRateLimitMap         map[string]uint64
	DataNodeReqZoneOpRateLimitMap       map[string]map[uint8]uint64
	DataNodeReqZoneVolOpRateLimitMap    map[string]map[string]map[uint8]uint64
	DataNodeReqVolPartRateLimitMap      map[string]uint64
	DataNodeReqVolOpPartRateLimitMap    map[string]map[uint8]uint64
	reqRateLimitMapMutex                sync.Mutex
	DataNodeDeleteLimitRate             uint64 //datanode delete limit rate
	DataNodeRepairTaskCount             uint64
	MetaNodeDeleteWorkerSleepMs         uint64
	ClientReadVolRateLimitMap           map[string]uint64
	ClientWriteVolRateLimitMap          map[string]uint64
	ClientVolOpRateLimitMap             map[string]map[uint8]int64
	ExtentMergeIno                      map[string][]uint64
	ExtentMergeSleepMs                  uint64
	peers                               []raftstore.PeerAddress
	peerAddrs                           []string
	heartbeatPort                       int64
	replicaPort                         int64
	diffSpaceUsage                      uint64
	DataPartitionsRecoverPoolSize       int32
	MetaPartitionsRecoverPoolSize       int32
}

func newClusterConfig() (cfg *clusterConfig) {
	cfg = new(clusterConfig)
	cfg.numberOfDataPartitionsToFree = defaultTobeFreedDataPartitionCount
	cfg.secondsToFreeDataPartitionAfterLoad = defaultSecondsToFreeDataPartitionAfterLoad
	cfg.NodeTimeOutSec = defaultNodeTimeOutSec
	cfg.MissingDataPartitionInterval = defaultMissingDataPartitionInterval
	cfg.DataPartitionTimeOutSec = defaultDataPartitionTimeOutSec
	cfg.IntervalToCheckDataPartition = defaultIntervalToCheckDataPartition
	cfg.IntervalToAlarmMissingDataPartition = defaultIntervalToAlarmMissingDataPartition
	cfg.numberOfDataPartitionsToLoad = defaultNumberOfDataPartitionsToLoad
	cfg.PeriodToLoadALLDataPartitions = defaultPeriodToLoadAllDataPartitions
	cfg.MetaNodeThreshold = defaultMetaPartitionMemUsageThreshold
	cfg.metaNodeReservedMem = defaultMetaNodeReservedMem
	cfg.diffSpaceUsage = defaultDiffSpaceUsage
	cfg.DataPartitionsRecoverPoolSize = defaultRecoverPoolSize
	cfg.MetaPartitionsRecoverPoolSize = defaultRecoverPoolSize
	cfg.MetaNodeReqOpRateLimitMap = make(map[uint8]uint64)
	cfg.DataNodeReqZoneRateLimitMap = make(map[string]uint64)
	cfg.DataNodeReqZoneOpRateLimitMap = make(map[string]map[uint8]uint64)
	cfg.DataNodeReqZoneVolOpRateLimitMap = make(map[string]map[string]map[uint8]uint64)
	cfg.DataNodeReqVolPartRateLimitMap = make(map[string]uint64)
	cfg.DataNodeReqVolOpPartRateLimitMap = make(map[string]map[uint8]uint64)
	cfg.ClientReadVolRateLimitMap = make(map[string]uint64)
	cfg.ClientWriteVolRateLimitMap = make(map[string]uint64)
	cfg.ClientVolOpRateLimitMap = make(map[string]map[uint8]int64)
	cfg.ExtentMergeIno = make(map[string][]uint64)
	return
}

func parsePeerAddr(peerAddr string) (id uint64, ip string, port uint64, err error) {
	peerStr := strings.Split(peerAddr, colonSplit)
	id, err = strconv.ParseUint(peerStr[0], 10, 64)
	if err != nil {
		return
	}
	port, err = strconv.ParseUint(peerStr[2], 10, 64)
	if err != nil {
		return
	}
	ip = peerStr[1]
	return
}

func (cfg *clusterConfig) parsePeers(peerStr string) error {
	peerArr := strings.Split(peerStr, commaSplit)
	cfg.peerAddrs = peerArr
	for _, peerAddr := range peerArr {
		id, ip, port, err := parsePeerAddr(peerAddr)
		if err != nil {
			return err
		}
		cfg.peers = append(cfg.peers, raftstore.PeerAddress{Peer: proto.Peer{ID: id}, Address: ip, HeartbeatPort: int(cfg.heartbeatPort), ReplicaPort: int(cfg.replicaPort)})
		address := fmt.Sprintf("%v:%v", ip, port)
		fmt.Println(address)
		AddrDatabase[id] = address
	}
	return nil
}
