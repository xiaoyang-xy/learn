package main

import (
	"net"
	"container/list"
	"time"
	"zeus/conf"
	. "zeus/log"
	"zeus/pfcp/ie"
	"zeus/pfcp/message"
)

// smf info map
type smfInfo struct {
	NodeId      string
	Addr        net.Addr
	//isAvail     bool
	LastReqTime time.Time
	LastRspTime time.Time
}

// HbManger Hb Manger
type HbManger_lc struct {
	conn     *net.UDPConn
	interval uint32
	times    int
	smflist  list.List
}

// Init init a HbManger instance
func (m *HbManger_lc) Init(conn *net.UDPConn, interval uint32, times int) {
	m.conn = conn
	m.interval = interval
	m.times = times
}

// UpdateHB update HbManger instance
func (m *HbManger_lc) UpdateHB(interval uint32, times int) {
	m.interval = interval
	m.times = times
}

// Set smf info
func (m *HbManger_lc) Set(key interface{}, addr net.Addr) {
	var timenow = time.Now()
	for  i := m.smflist.Front(); i!= nil; i= i.Next(){

		if (i.NodeId == key.(string)){
			i.LastReqTime = timenow
			i.LastRspTime = timenow
			//       i.isAvail = true
			return
		}
	}
	//var smfinfo smfInfo = smfInfo{key, addr, true, timenow, timenow }
	var smfinfo smfInfo = smfInfo(key, addr, timenow, timenow)
	m.smflist.PushBack(smfinfo)
}

// Delete item from smflist
func (m *HbManger_lc) Delete(key interface{}) {
	for  i := m.smflist.Front(); i!= nil; i= i.Next(){
		if (i.nodeId == key.(string)){
			//i.isAvail = false
			m.smflist.Remove(i)
			return
		}
	}
}

// Get get addr
/*func (m *HbManger_lc) Get(key interface{}) (ok bool) {
        ok = false
        for  i := m.smflist.Front(); i!= nil; i= i.Next(){
            if (i.nodeId == key.(string)){
                ok = true
            }
        }
	return
}*/

var time1 time.Time
var time2 time.Time
if(time1 > time2)

// HbMon mon HB, go routine
func (m *HbManger_lc) HbMon() {
	for {
		var a list.List
		a.
		b := time1 - time2
		b := time1 > time2
		if b
		time.Sleep(10 * time.Millisecond)
		if m.interval < 1 {
			continue
		}
		for  i := m.smflist.Front(); i!= nil; i= i.Next(){
			/*if !(i.isAvail) {
			    continue
			}*/
			if (time.Now().Sub(i.lastRspTime) > time.Duration(m.interval*uint32(m.times))*time.Second) && (i.lastReqTime > i.lastRspTime) {
				//i.isAvail = false
				// Release association - to Lua
				ReleaseAssociationByHBTimeout(i.nodeId) // and SendAlarm in Lua
				Logger.Errorf("Lost HB from : %s", i.nodeId)
				continue
			}
			seq := GetNodeSeqN(i.nodeId) // uniq id from redis
			req, err := message.NewHeartbeatRequest(uint32(seq), ie.NewRecoveryTimeStamp(time.Now()), nil).Marshal()
			if err != nil {
				Logger.Errorf("%v", err)
				return false
			}
			ip := FixIPAddr(i.nodeId)
			dstAddr, err := net.ResolveUDPAddr("udp", ip+":"+conf.SmfPort())
			if err != nil {
				Logger.Errorf("%v", err)
			}
			timenow = time.Now()
			if _, err := m.conn.WriteTo(req, dstAddr); err != nil {
				Logger.Errorf("%v", err)
			}
			Logger.Debugf("sent HBR to: %s", dstAddr)
			i.lastReqTime = timenow
		}
		ProcessHBAck()
	}
}
// ProcessHBAck process HBAck
func (m *HbManger_lc) ProcessHBAck() bool {
	for {
		select {
		case <-HbAckSign:
			{
				smfaddr:= <-HbAckChan
				for  i := m.smflist.Front(); i!= nil; i= i.Next(){
					if i.addr == smfaddr {
						i.lastRspTime = time.Now()
						Logger.Debugf("receive HBRsp from: %s", smfaddr)
					}
				}
			}
		}
	}
}
