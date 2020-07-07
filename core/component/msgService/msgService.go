package msgService

import (
	"log"
	"reflect"
	"sync"
)

type MsgInfo struct {
	//Route int
	ReqType  reflect.Type
	RespType reflect.Type
}

type msgMgr struct {
	msgMap map[int]*MsgInfo
}

var (
	msgService *msgMgr
	once       sync.Once
)

func GetMsgService() *msgMgr {
	once.Do(func() {
		//msgService = &msgMgr{msgMap:map[int]*MsgInfo{}}
		msgService = new(msgMgr)
		msgService.msgMap = make(map[int]*MsgInfo)
	})
	return msgService
}

func (m *msgMgr) Register(route int, msgReq interface{}, msgResp interface{}) {
	reqType := reflect.TypeOf(msgReq)
	if reqType == nil || reqType.Kind() != reflect.Ptr {
		log.Fatal("message request pointer required")
		return
	}
	respType := reflect.TypeOf(msgResp)
	if respType == nil || respType.Kind() != reflect.Ptr {
		log.Fatal("message response pointer required")
	}
	if _, ok := m.msgMap[route]; ok {
		//log.Fatal("route %s is already registered", route)
		return
	}

	i := new(MsgInfo)
	//i.Route = method
	i.ReqType = reqType
	i.RespType = respType
	m.msgMap[route] = i
}

func (m *msgMgr) GetMsgByRouteId(route int) (info *MsgInfo, ok bool) {
	info, ok = m.msgMap[route]
	return
}
