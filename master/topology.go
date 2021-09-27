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
	"github.com/chubaofs/chubaofs/proto"
	"github.com/chubaofs/chubaofs/util/errors"
	"github.com/chubaofs/chubaofs/util/log"
	"sort"
	"strings"
	"sync"
)

type topology struct {
	dataNodes            *sync.Map
	metaNodes            *sync.Map
	zoneMap              *sync.Map
	zoneIndexForDataNode int
	zoneIndexForMetaNode int
	zones                []*Zone
	zoneLock             sync.RWMutex
}

func newTopology() (t *topology) {
	t = new(topology)
	t.zoneMap = new(sync.Map)
	t.dataNodes = new(sync.Map)
	t.metaNodes = new(sync.Map)
	t.zones = make([]*Zone, 0)
	return
}

func (t *topology) zoneLen() int {
	t.zoneLock.RLock()
	defer t.zoneLock.RUnlock()
	return len(t.zones)
}

func (t *topology) clear() {
	t.dataNodes.Range(func(key, value interface{}) bool {
		t.dataNodes.Delete(key)
		return true
	})
	t.metaNodes.Range(func(key, value interface{}) bool {
		t.metaNodes.Delete(key)
		return true
	})
}

func (t *topology) putZone(zone *Zone) (err error) {
	t.zoneLock.Lock()
	defer t.zoneLock.Unlock()
	if _, ok := t.zoneMap.Load(zone.name); ok {
		return fmt.Errorf("zone[%v] has exist", zone.name)
	}
	t.zoneMap.Store(zone.name, zone)
	t.zones = append(t.zones, zone)
	return
}

func (t *topology) putZoneIfAbsent(zone *Zone) (beStoredZone *Zone) {
	t.zoneLock.Lock()
	defer t.zoneLock.Unlock()
	oldZone, ok := t.zoneMap.Load(zone.name)
	if ok {
		return oldZone.(*Zone)
	}
	t.zoneMap.Store(zone.name, zone)
	t.zones = append(t.zones, zone)
	beStoredZone = zone
	return
}

func (t *topology) getZone(name string) (zone *Zone, err error) {
	if name == "" {
		return nil, fmt.Errorf("zone name is empty")
	}
	t.zoneMap.Range(func(zoneName, value interface{}) bool {
		if zoneName != name {
			return true
		}
		zone = value.(*Zone)
		return true
	})
	if zone == nil {
		return nil, fmt.Errorf("zone[%v] is not found", name)
	}
	return
}

func (t *topology) putDataNode(dataNode *DataNode) (err error) {

	if _, ok := t.dataNodes.Load(dataNode.Addr); ok {
		return
	}
	zone, err := t.getZone(dataNode.ZoneName)
	if err != nil {
		return
	}

	zone.putDataNode(dataNode)
	t.putDataNodeToCache(dataNode)
	return
}

func (t *topology) putDataNodeToCache(dataNode *DataNode) {
	t.dataNodes.Store(dataNode.Addr, dataNode)
}

func (t *topology) deleteDataNode(dataNode *DataNode) {
	zone, err := t.getZone(dataNode.ZoneName)
	if err != nil {
		return
	}
	zone.deleteDataNode(dataNode)
	t.dataNodes.Delete(dataNode.Addr)
}

func (t *topology) getZoneByDataNode(dataNode *DataNode) (zone *Zone, err error) {
	_, ok := t.dataNodes.Load(dataNode.Addr)
	if !ok {
		return nil, errors.Trace(dataNodeNotFound(dataNode.Addr), "%v not found", dataNode.Addr)
	}

	return t.getZone(dataNode.ZoneName)
}

func (t *topology) getZoneByMetaNode(metaNode *MetaNode) (zone *Zone, err error) {
	_, ok := t.metaNodes.Load(metaNode.Addr)
	if !ok {
		return nil, errors.Trace(metaNodeNotFound(metaNode.Addr), "%v not found", metaNode.Addr)
	}

	return t.getZone(metaNode.ZoneName)
}

func (t *topology) putMetaNode(metaNode *MetaNode) (err error) {
	if _, ok := t.metaNodes.Load(metaNode.Addr); ok {
		return
	}
	zone, err := t.getZone(metaNode.ZoneName)
	if err != nil {
		return
	}
	zone.putMetaNode(metaNode)
	t.putMetaNodeToCache(metaNode)
	return
}

func (t *topology) deleteMetaNode(metaNode *MetaNode) {
	t.metaNodes.Delete(metaNode.Addr)
	zone, err := t.getZone(metaNode.ZoneName)
	if err != nil {
		return
	}
	zone.deleteMetaNode(metaNode)
}

func (t *topology) putMetaNodeToCache(metaNode *MetaNode) {
	t.metaNodes.Store(metaNode.Addr, metaNode)
}

type nodeSetCollection []*nodeSet

func (nsc nodeSetCollection) Len() int {
	return len(nsc)
}

func (nsc nodeSetCollection) Less(i, j int) bool {
	return nsc[i].metaNodeLen() < nsc[j].metaNodeLen()
}

func (nsc nodeSetCollection) Swap(i, j int) {
	nsc[i], nsc[j] = nsc[j], nsc[i]
}

type nodeSetCollectionForDataNode []*nodeSet

func (nsc nodeSetCollectionForDataNode) Len() int {
	return len(nsc)
}

func (nsc nodeSetCollectionForDataNode) Less(i, j int) bool {
	return nsc[i].dataNodeLen() < nsc[j].dataNodeLen()
}

func (nsc nodeSetCollectionForDataNode) Swap(i, j int) {
	nsc[i], nsc[j] = nsc[j], nsc[i]
}

type nodeSet struct {
	ID        uint64
	Capacity  int
	zoneName  string
	metaNodes *sync.Map
	dataNodes *sync.Map
	sync.RWMutex
}

func newNodeSet(id uint64, cap int, zoneName string) *nodeSet {
	ns := &nodeSet{
		ID:        id,
		Capacity:  cap,
		zoneName:  zoneName,
		metaNodes: new(sync.Map),
		dataNodes: new(sync.Map),
	}
	return ns
}

func (ns *nodeSet) metaNodeLen() (count int) {
	ns.RLock()
	defer ns.RUnlock()
	ns.metaNodes.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return
}

func (ns *nodeSet) dataNodeLen() (count int) {
	ns.RLock()
	defer ns.RUnlock()
	ns.dataNodes.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return
}

func (ns *nodeSet) putMetaNode(metaNode *MetaNode) {
	ns.metaNodes.Store(metaNode.Addr, metaNode)
}

func (ns *nodeSet) deleteMetaNode(metaNode *MetaNode) {
	ns.metaNodes.Delete(metaNode.Addr)
}

// can Write For DataNode With Exclude Hosts
func (ns *nodeSet) canWriteForDataNode(excludeHosts []string, replicaNum int) bool {
	var count int
	ns.dataNodes.Range(func(key, value interface{}) bool {
		node := value.(*DataNode)
		if contains(excludeHosts, node.Addr) == true {
			log.LogDebugf("contains return")
			return true
		}
		if node.isWriteAble() {
			count++
		}
		if count >= replicaNum {
			return false
		}
		return true
	})
	log.LogInfof("canWriteForDataNode zone[%v], ns[%v],count[%v], replicaNum[%v]",
		ns.zoneName, ns.ID, count, replicaNum)
	return count >= replicaNum
}

func (ns *nodeSet) canWriteForMetaNode(replicaNum int) bool {
	var count int
	ns.metaNodes.Range(func(key, value interface{}) bool {
		node := value.(*MetaNode)
		if node.isMixedMetaNode() {
			return true
		}
		if node.isWritable() {
			count++
		}
		if count >= replicaNum {
			return false
		}
		return true
	})
	log.LogInfof("canWriteForMetaNode zone[%v], ns[%v],count[%v] replicaNum[%v]",
		ns.zoneName, ns.ID, count, replicaNum)
	return count >= replicaNum
}

func (ns *nodeSet) putDataNode(dataNode *DataNode) {
	ns.dataNodes.Store(dataNode.Addr, dataNode)
}

func (ns *nodeSet) deleteDataNode(dataNode *DataNode) {
	ns.dataNodes.Delete(dataNode.Addr)
}

func (t *topology) isSingleZone() bool {
	t.zoneLock.RLock()
	defer t.zoneLock.RUnlock()
	var zoneLen int
	t.zoneMap.Range(func(zoneName, value interface{}) bool {
		zoneLen++
		return true
	})
	return zoneLen == 1
}

func (t *topology) getAllZones() (zones []*Zone) {
	t.zoneLock.RLock()
	defer t.zoneLock.RUnlock()
	zones = make([]*Zone, 0)
	t.zoneMap.Range(func(zoneName, value interface{}) bool {
		zone := value.(*Zone)
		zones = append(zones, zone)
		return true
	})
	return
}

func (t *topology) getZoneByIndex(index int) (zone *Zone) {
	t.zoneLock.RLock()
	defer t.zoneLock.RUnlock()
	return t.zones[index]
}

func calculateDemandWriteNodes(zoneNum, replicaNum int) (demandWriteNodes int) {
	if zoneNum == 1 {
		demandWriteNodes = replicaNum
	}
	if zoneNum >= 2 {
		demandWriteNodes = 2
	}
	if demandWriteNodes > replicaNum {
		demandWriteNodes = 1
	}
	return
}

func (t *topology) allocZonesForMetaNode(zoneName string, replicaNum int, excludeZone []string) (candidateZones []*Zone, err error) {
	var initCandidateZones []*Zone
	initCandidateZones = make([]*Zone, 0)
	zoneList := strings.Split(zoneName, ",")
	if t.isSingleZone() {
		return t.getAllZones(), nil
	}
	if excludeZone == nil {
		excludeZone = make([]string, 0)
	}
	for _, z := range zoneList {
		var zone *Zone
		if zone, err = t.getZone(z); err != nil {
			return
		}
		initCandidateZones = append(initCandidateZones, zone)
	}
	demandWriteNodes := calculateDemandWriteNodes(len(zoneList), replicaNum)
	candidateZones = make([]*Zone, 0)
	for _, zone := range initCandidateZones {
		if zone.status == unavailableZone {
			continue
		}
		if zone.canWriteForMetaNode(uint8(demandWriteNodes)) {
			candidateZones = append(candidateZones, zone)
		}
		if len(candidateZones) >= len(zoneList) {
			break
		}
	}
	//if there is no space in the zone for single zone partition, randomly choose another zone
	if len(candidateZones) < 1 && len(zoneList) == 1 {
		initCandidateZones = t.getAllZones()
		for _, zone := range initCandidateZones {
			if zone.status == unavailableZone {
				continue
			}
			if zone.canWriteForMetaNode(uint8(demandWriteNodes)) {
				candidateZones = append(candidateZones, zone)
			}
		}
	}

	//if across zone,candidateZones must be larger than or equal with 2,otherwise,must have a candidate zone
	if (replicaNum == 3 && len(zoneList) >= 2 && len(candidateZones) < 2) || len(candidateZones) < 1 {
		log.LogError(fmt.Sprintf("action[allocZonesForMetaNode],reqZoneNum[%v],candidateZones[%v],demandWriteNodes[%v],err:%v",
			len(zoneList), len(candidateZones), demandWriteNodes, proto.ErrNoZoneToCreateMetaPartition))
		return nil, proto.ErrNoZoneToCreateMetaPartition
	}
	err = nil
	return
}

//allocate zones according to the specified zoneName and replicaNum
func (t *topology) allocZonesForDataNode(zoneName string, replicaNum int, excludeZone []string) (candidateZones []*Zone, err error) {
	var initCandidateZones []*Zone
	initCandidateZones = make([]*Zone, 0)
	zoneList := strings.Split(zoneName, ",")
	if t.isSingleZone() {
		return t.getAllZones(), nil
	}
	for _, z := range zoneList {
		var zone *Zone
		if zone, err = t.getZone(z); err != nil {
			return
		}
		initCandidateZones = append(initCandidateZones, zone)
	}
	demandWriteNodes := calculateDemandWriteNodes(len(zoneList), replicaNum)
	candidateZones = make([]*Zone, 0)
	for _, zone := range initCandidateZones {
		if zone.status == unavailableZone {
			continue
		}
		if contains(excludeZone, zone.name) {
			continue
		}
		if zone.canWriteForDataNode(uint8(demandWriteNodes)) {
			candidateZones = append(candidateZones, zone)
		}
		if len(candidateZones) >= len(zoneList) {
			break
		}
	}
	//if there is no space in the zone for single zone partition, randomly choose a zone from all zones
	if len(candidateZones) < 1 && len(zoneList) == 1 {
		initCandidateZones = t.getAllZones()
		for _, zone := range initCandidateZones {
			if zone.status == unavailableZone {
				continue
			}
			if contains(excludeZone, zone.name) {
				continue
			}
			if zone.canWriteForDataNode(uint8(demandWriteNodes)) {
				candidateZones = append(candidateZones, zone)
			}
		}
	}
	//if across zone,candidateZones must be larger than or equal with 2, if not across zone, must have one candidate zone
	if (replicaNum == 3 && len(zoneList) >= 2 && len(candidateZones) < 2) || len(candidateZones) < 1 {
		log.LogError(fmt.Sprintf("action[allocZonesForDataNode],reqZoneNum[%v],candidateZones[%v],demandWriteNodes[%v],err:%v",
			len(zoneList), len(candidateZones), demandWriteNodes, proto.ErrNoZoneToCreateDataPartition))
		return nil, errors.NewError(proto.ErrNoZoneToCreateDataPartition)
	}
	err = nil
	return
}

func (ns *nodeSet) dataNodeCount() int {
	var count int
	ns.dataNodes.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (ns *nodeSet) getAvailDataNodeHosts(excludeHosts []string, replicaNum int) (hosts []string, peers []proto.Peer, err error) {
	return getAvailHosts(ns.dataNodes, excludeHosts, replicaNum, selectDataNode)
}

func (ns *nodeSet) metaNodeCount() int {
	var count int
	ns.metaNodes.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// Zone stores all the zone related information
type Zone struct {
	name                string
	setIndexForDataNode int
	setIndexForMetaNode int
	status              int
	dataNodes           *sync.Map
	metaNodes           *sync.Map
	nodeSetMap          map[uint64]*nodeSet
	nsLock              sync.RWMutex
	sync.RWMutex
}

func newZone(name string) (zone *Zone) {
	zone = &Zone{name: name}
	zone.status = normalZone
	zone.dataNodes = new(sync.Map)
	zone.metaNodes = new(sync.Map)
	zone.nodeSetMap = make(map[uint64]*nodeSet)
	return
}

func (zone *Zone) setStatus(status int) {
	zone.status = status
}

func (zone *Zone) getStatus() int {
	return zone.status
}

func (zone *Zone) getStatusToString() string {
	if zone.status == normalZone {
		return "available"
	} else {
		return "unavailable"
	}
}

func (zone *Zone) isSingleNodeSet() bool {
	zone.RLock()
	defer zone.RUnlock()
	return len(zone.nodeSetMap) == 1
}

func (zone *Zone) getNodeSet(setID uint64) (ns *nodeSet, err error) {
	zone.nsLock.RLock()
	defer zone.nsLock.RUnlock()
	ns, ok := zone.nodeSetMap[setID]
	if !ok {
		return nil, errors.NewErrorf("set %v not found", setID)
	}
	return
}

func (zone *Zone) putNodeSet(ns *nodeSet) (err error) {
	zone.nsLock.Lock()
	defer zone.nsLock.Unlock()
	if _, ok := zone.nodeSetMap[ns.ID]; ok {
		return fmt.Errorf("nodeSet [%v] has exist", ns.ID)
	}
	zone.nodeSetMap[ns.ID] = ns
	return
}

func (zone *Zone) createNodeSet(c *Cluster) (ns *nodeSet, err error) {
	id, err := c.idAlloc.allocateCommonID()
	if err != nil {
		return
	}
	ns = newNodeSet(id, c.cfg.nodeSetCapacity, zone.name)
	if err = c.syncAddNodeSet(ns); err != nil {
		return
	}
	if err = zone.putNodeSet(ns); err != nil {
		return
	}
	return
}

func (zone *Zone) getAllNodeSet() (nsc nodeSetCollection) {
	zone.nsLock.RLock()
	defer zone.nsLock.RUnlock()
	nsc = make(nodeSetCollection, 0)
	for _, ns := range zone.nodeSetMap {
		nsc = append(nsc, ns)
	}
	return
}

func (zone *Zone) getAllNodeSetForDataNode() (nsc nodeSetCollectionForDataNode) {
	zone.nsLock.RLock()
	defer zone.nsLock.RUnlock()
	nsc = make(nodeSetCollectionForDataNode, 0)
	for _, ns := range zone.nodeSetMap {
		nsc = append(nsc, ns)
	}
	return
}

func (zone *Zone) getAvailNodeSetForMetaNode() (nset *nodeSet) {
	allNodeSet := zone.getAllNodeSet()
	sort.Sort(sort.Reverse(allNodeSet))
	for _, ns := range allNodeSet {
		if ns.metaNodeLen() < ns.Capacity {
			nset = ns
			return
		}
	}
	return
}

func (zone *Zone) getAvailNodeSetForDataNode() (nset *nodeSet) {
	allNodeSet := zone.getAllNodeSetForDataNode()
	sort.Sort(sort.Reverse(allNodeSet))
	for _, ns := range allNodeSet {
		if ns.dataNodeLen() < ns.Capacity {
			nset = ns
			return
		}
	}
	return
}

func (zone *Zone) putDataNode(dataNode *DataNode) (err error) {
	var ns *nodeSet
	if ns, err = zone.getNodeSet(dataNode.NodeSetID); err != nil {
		log.LogErrorf("action[putDataNode] nodeSet[%v] not found", dataNode.NodeSetID)
		return
	}
	ns.putDataNode(dataNode)
	zone.dataNodes.Store(dataNode.Addr, dataNode)
	return
}

func (zone *Zone) getDataNode(addr string) (dataNode *DataNode, err error) {
	value, ok := zone.dataNodes.Load(addr)
	if !ok {
		return nil, errors.Trace(dataNodeNotFound(addr), "%v not found", addr)
	}
	dataNode = value.(*DataNode)
	return
}
func (zone *Zone) deleteDataNode(dataNode *DataNode) {
	ns, err := zone.getNodeSet(dataNode.NodeSetID)
	if err != nil {
		log.LogErrorf("action[zoneDeleteDataNode] nodeSet[%v] not found", dataNode.NodeSetID)
		return
	}
	ns.deleteDataNode(dataNode)
	zone.dataNodes.Delete(dataNode.Addr)
}

func (zone *Zone) putMetaNode(metaNode *MetaNode) (err error) {
	var ns *nodeSet
	if ns, err = zone.getNodeSet(metaNode.NodeSetID); err != nil {
		log.LogErrorf("action[zonePutMetaNode] nodeSet[%v] not found", metaNode.NodeSetID)
		return
	}
	ns.putMetaNode(metaNode)
	zone.metaNodes.Store(metaNode.Addr, metaNode)
	return
}

func (zone *Zone) getMetaNode(addr string) (metaNode *MetaNode, err error) {
	value, ok := zone.metaNodes.Load(addr)
	if !ok {
		return nil, errors.Trace(metaNodeNotFound(addr), "%v not found", addr)
	}
	metaNode = value.(*MetaNode)
	return
}


func (zone *Zone) deleteMetaNode(metaNode *MetaNode) (err error) {
	ns, err := zone.getNodeSet(metaNode.NodeSetID)
	if err != nil {
		log.LogErrorf("action[zoneDeleteMetaNode] nodeSet[%v] not found", metaNode.NodeSetID)
		return
	}
	ns.deleteMetaNode(metaNode)
	zone.metaNodes.Delete(metaNode.Addr)
	return
}

func (zone *Zone) allocNodeSetForDataNode(excludeNodeSets []uint64, excludeHosts []string, replicaNum uint8) (ns *nodeSet, err error) {
	nset := zone.getAllNodeSet()
	if nset == nil {
		return nil, errors.NewError(proto.ErrNoNodeSetToCreateDataPartition)
	}
	zone.nsLock.Lock()
	defer zone.nsLock.Unlock()
	for i := 0; i < len(nset); i++ {
		if zone.setIndexForDataNode >= len(nset) {
			zone.setIndexForDataNode = 0
		}
		ns = nset[zone.setIndexForDataNode]
		zone.setIndexForDataNode++
		if containsID(excludeNodeSets, ns.ID) {
			continue
		}
		if ns.canWriteForDataNode(excludeHosts, int(replicaNum)) {
			return
		}
	}
	log.LogErrorf("action[allocNodeSetForDataNode],nset len[%v],excludeNodeSets[%v],rNum[%v] err:%v",
		nset.Len(), excludeNodeSets, replicaNum, proto.ErrNoNodeSetToCreateDataPartition)
	return nil, errors.NewError(proto.ErrNoNodeSetToCreateDataPartition)
}

func (zone *Zone) allocNodeSetForMetaNode(excludeNodeSets []uint64, replicaNum uint8) (ns *nodeSet, err error) {
	nset := zone.getAllNodeSet()
	if nset == nil {
		return nil, proto.ErrNoNodeSetToCreateMetaPartition
	}
	zone.nsLock.Lock()
	defer zone.nsLock.Unlock()
	for i := 0; i < len(nset); i++ {
		if zone.setIndexForMetaNode >= len(nset) {
			zone.setIndexForMetaNode = 0
		}
		ns = nset[zone.setIndexForMetaNode]
		zone.setIndexForMetaNode++
		if containsID(excludeNodeSets, ns.ID) {
			continue
		}
		if ns.canWriteForMetaNode(int(replicaNum)) {
			return
		}
	}
	log.LogError(fmt.Sprintf("action[allocNodeSetForMetaNode],zone[%v],excludeNodeSets[%v],rNum[%v],err:%v",
		zone.name, excludeNodeSets, replicaNum, proto.ErrNoNodeSetToCreateMetaPartition))
	return nil, proto.ErrNoNodeSetToCreateMetaPartition
}

func (zone *Zone) canWriteForDataNode(replicaNum uint8) (can bool) {
	zone.RLock()
	defer zone.RUnlock()
	var leastAlive uint8
	zone.dataNodes.Range(func(addr, value interface{}) bool {
		dataNode := value.(*DataNode)
		if dataNode.isActive == true && dataNode.isWriteAble() == true {
			leastAlive++
		}
		if leastAlive >= replicaNum {
			can = true
			return false
		}
		return true
	})
	fmt.Printf("canWriteForDataNode leastAlive[%v],replicaNum[%v],count[%v]\n", leastAlive, replicaNum, zone.dataNodeCount())
	return
}

func (zone *Zone) canWriteForMetaNode(replicaNum uint8) (can bool) {
	zone.RLock()
	defer zone.RUnlock()
	var leastAlive uint8
	zone.metaNodes.Range(func(addr, value interface{}) bool {
		metaNode := value.(*MetaNode)
		if metaNode.isMixedMetaNode() {
			return true
		}
		if metaNode.IsActive == true && metaNode.isWritable() == true {
			leastAlive++
		}
		if leastAlive >= replicaNum {
			can = true
			return false
		}
		return true
	})
	return
}

func (zone *Zone) getDataNodeMaxTotal() (maxTotal uint64) {
	zone.dataNodes.Range(func(key, value interface{}) bool {
		dataNode := value.(*DataNode)
		if dataNode.Total > maxTotal {
			maxTotal = dataNode.Total
		}
		return true
	})
	return
}

func (zone *Zone) getAvailDataNodeHosts(excludeNodeSets []uint64, excludeHosts []string, replicaNum int) (newHosts []string, peers []proto.Peer, err error) {
	if replicaNum == 0 {
		return
	}
	ns, err := zone.allocNodeSetForDataNode(excludeNodeSets, excludeHosts, uint8(replicaNum))
	if err != nil {
		return nil, nil, errors.Trace(err, "zone[%v] alloc node set,replicaNum[%v]", zone.name, replicaNum)
	}
	return ns.getAvailDataNodeHosts(excludeHosts, replicaNum)
}

func (zone *Zone) getAvailMetaNodeHosts(excludeNodeSets []uint64, excludeHosts []string, replicaNum int) (newHosts []string, peers []proto.Peer, err error) {
	if replicaNum == 0 {
		return
	}
	ns, err := zone.allocNodeSetForMetaNode(excludeNodeSets, uint8(replicaNum))
	if err != nil {
		return nil, nil, errors.NewErrorf("zone[%v],err[%v]", zone.name, err)
	}
	return ns.getAvailMetaNodeHosts(excludeHosts, replicaNum)

}

func (zone *Zone) dataNodeCount() (len int) {

	zone.dataNodes.Range(func(key, value interface{}) bool {
		len++
		return true
	})
	return
}

// 1.get all node set with avail space then sort them
// 2.if only get two node set and the sum of their nodes more than 18(c.cfg.nodeSetCapacity),
// if less than min count or the difference between the number of their nodes large than defaultMinusOfNodeSetCount,average distribution them
// 3.fill large mode set by merging node from small node set to large node set
// 4.if there is one node set less than c.cfg.nodeSetCapacity/2 -1, get the largest node set and average distribution them
func (zone *Zone) mergeNodeSetForMetaNode(c *Cluster) {
	if zone.isSingleNodeSet() {
		return
	}
	nsc := make(nodeSetCollection, 0)
	zone.nsLock.RLock()
	for _, ns := range zone.nodeSetMap {
		if ns.metaNodeLen() == 0 || ns.metaNodeLen() >= ns.Capacity {
			continue
		}
		nsc = append(nsc, ns)
	}
	zone.nsLock.RUnlock()
	sort.Sort(nsc)
	if len(nsc) == 0 {
		return
	}
	minNodeSetCount := c.cfg.nodeSetCapacity/2 - 1
	if len(nsc) == 2 && nsc[0].metaNodeLen()+nsc[1].metaNodeLen() > c.cfg.nodeSetCapacity {
		minusNodeCount := nsc[1].metaNodeLen() - nsc[0].metaNodeLen()
		if nsc[0].metaNodeLen() < minNodeSetCount || minusNodeCount > defaultMinusOfNodeSetCount {
			// Average distribution, merge node from large to small
			avgCount := (nsc[0].metaNodeLen() + nsc[1].metaNodeLen()) / 2
			successNum, err := zone.batchMergeNodeSetForMetaNode(c, avgCount-nsc[0].metaNodeLen(), nsc[1].ID, nsc[0].ID)
			if err != nil {
				log.LogErrorf("action[mergeNodeSetForMetaNode] repair two ns sourceID[%v], targetID[%v] success num[%v] err[%v]", nsc[1].ID, nsc[0].ID, successNum, err)
				return
			}
			log.LogWarnf("action[mergeNodeSetForMetaNode] repair two ns sourceID[%v], targetID[%v] success num[%v] ", nsc[1].ID, nsc[0].ID, successNum)
		}
		return
	}

	small := 0
	large := len(nsc) - 1
	for small < large {
		if nsc[small].metaNodeLen() == 0 {
			small++
			continue
		}
		if nsc[large].metaNodeLen() >= nsc[large].Capacity {
			large--
			continue
		}
		count := nsc[small].metaNodeLen()
		lackCount := nsc[large].Capacity - nsc[large].metaNodeLen()
		if count > lackCount {
			count = lackCount
		}
		successNum, err := zone.batchMergeNodeSetForMetaNode(c, count, nsc[small].ID, nsc[large].ID)
		if err != nil {
			log.LogErrorf("action[mergeNodeSetForMetaNode] fill large ns sourceID[%v] targetID[%v] success num[%v] err[%v]", nsc[small].ID, nsc[large].ID, successNum, err)
			return
		}
		log.LogWarnf("action[mergeNodeSetForMetaNode] fill large ns sourceID[%v] targetID[%v] success num[%v]", nsc[small].ID, nsc[large].ID, successNum)
	}

	if nsc[small].metaNodeLen() < minNodeSetCount && nsc[small].metaNodeLen() != 0 {
		allNodeSet := zone.getAllNodeSet()
		sort.Sort(sort.Reverse(allNodeSet))
		if len(allNodeSet) == 0 {
			return
		}
		largestNodeSet := allNodeSet[0]
		// if can be merged to one node set or node set is same, return
		if largestNodeSet.ID == nsc[small].ID || largestNodeSet.metaNodeLen()+nsc[small].metaNodeLen() <= largestNodeSet.Capacity {
			return
		}
		avgCount := (nsc[small].metaNodeLen() + largestNodeSet.metaNodeLen()) / 2
		successNum, err := zone.batchMergeNodeSetForMetaNode(c, avgCount-nsc[small].metaNodeLen(), largestNodeSet.ID, nsc[small].ID)
		if err != nil {
			log.LogErrorf("action[mergeNodeSetForMetaNode] repair too small ns sourceID[%v], targetID[%v] success num[%v] err[%v]", largestNodeSet.ID, nsc[small].ID, successNum, err)
			return
		}
		log.LogWarnf("action[mergeNodeSetForMetaNode] repair too small ns sourceID[%v], targetID[%v] success num[%v] ", largestNodeSet.ID, nsc[small].ID, successNum)
	}
	return
}

func (zone *Zone) mergeNodeSetForDataNode(c *Cluster) {
	if zone.isSingleNodeSet() {
		return
	}
	nsc := make(nodeSetCollectionForDataNode, 0)
	zone.nsLock.RLock()
	for _, ns := range zone.nodeSetMap {
		if ns.dataNodeLen() == 0 || ns.dataNodeLen() >= ns.Capacity {
			continue
		}
		nsc = append(nsc, ns)
	}
	zone.nsLock.RUnlock()
	sort.Sort(nsc)
	if len(nsc) == 0 {
		return
	}
	minNodeSetCount := c.cfg.nodeSetCapacity/2 - 1
	if len(nsc) == 2 && nsc[0].dataNodeLen()+nsc[1].dataNodeLen() > c.cfg.nodeSetCapacity {
		minusNodeCount := nsc[1].dataNodeLen() - nsc[0].dataNodeLen()
		if nsc[0].dataNodeLen() < minNodeSetCount || minusNodeCount > defaultMinusOfNodeSetCount {
			// Average distribution, merge node from large to small
			avgCount := (nsc[0].dataNodeLen() + nsc[1].dataNodeLen()) / 2
			successNum, err := zone.batchMergeNodeSetForDataNode(c, avgCount-nsc[0].dataNodeLen(), nsc[1].ID, nsc[0].ID)
			if err != nil {
				log.LogErrorf("action[mergeNodeSetForDataNode] repair two ns sourceID[%v], targetID[%v] success num[%v] err[%v]", nsc[1].ID, nsc[0].ID, successNum, err)
				return
			}
			log.LogWarnf("action[mergeNodeSetForDataNode] repair two ns sourceID[%v], targetID[%v] success num[%v] ", nsc[1].ID, nsc[0].ID, successNum)
		}
		return
	}

	small := 0
	large := len(nsc) - 1
	for small < large {
		if nsc[small].dataNodeLen() == 0 {
			small++
			continue
		}
		if nsc[large].dataNodeLen() >= nsc[large].Capacity {
			large--
			continue
		}
		count := nsc[small].dataNodeLen()
		lackCount := nsc[large].Capacity - nsc[large].dataNodeLen()
		if count > lackCount {
			count = lackCount
		}
		successNum, err := zone.batchMergeNodeSetForDataNode(c, count, nsc[small].ID, nsc[large].ID)
		if err != nil {
			log.LogErrorf("action[mergeNodeSetForDataNode] fill large ns sourceID[%v] targetID[%v] success num[%v] err[%v]", nsc[small].ID, nsc[large].ID, successNum, err)
			return
		}
		log.LogWarnf("action[mergeNodeSetForDataNode] fill large ns sourceID[%v] targetID[%v] success num[%v]", nsc[small].ID, nsc[large].ID, successNum)
	}

	if nsc[small].dataNodeLen() < minNodeSetCount && nsc[small].dataNodeLen() != 0 {
		allNodeSet := zone.getAllNodeSetForDataNode()
		sort.Sort(sort.Reverse(allNodeSet))
		if len(allNodeSet) == 0 {
			return
		}
		largestNodeSet := allNodeSet[0]
		// if can be merged to one node set or node set is same, return
		if largestNodeSet.ID == nsc[small].ID || largestNodeSet.dataNodeLen()+nsc[small].dataNodeLen() <= largestNodeSet.Capacity {
			return
		}
		avgCount := (nsc[small].dataNodeLen() + largestNodeSet.dataNodeLen()) / 2
		successNum, err := zone.batchMergeNodeSetForDataNode(c, avgCount-nsc[small].dataNodeLen(), largestNodeSet.ID, nsc[small].ID)
		if err != nil {
			log.LogErrorf("action[mergeNodeSetForDataNode] repair too small ns sourceID[%v], targetID[%v] success num[%v] err[%v]", largestNodeSet.ID, nsc[small].ID, successNum, err)
			return
		}
		log.LogWarnf("action[mergeNodeSetForDataNode] repair too small ns sourceID[%v], targetID[%v] success num[%v] ", largestNodeSet.ID, nsc[small].ID, successNum)
	}
	return
}

func (zone *Zone) batchMergeNodeSetForMetaNode(c *Cluster, count int, sourceID, targetID uint64) (successNum int, err error) {
	for i := 0; i < count; i++ {
		if err = c.adjustNodeSetForMetaNode(zone.name, "", sourceID, targetID); err != nil {
			return
		}
		successNum++
	}
	return
}

func (zone *Zone) batchMergeNodeSetForDataNode(c *Cluster, count int, sourceID, targetID uint64) (successNum int, err error) {
	for i := 0; i < count; i++ {
		if err = c.adjustNodeSetForDataNode(zone.name, "", sourceID, targetID); err != nil {
			return
		}
		successNum++
	}
	return
}
