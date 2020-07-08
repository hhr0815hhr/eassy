package gateway

import (
	"eassy/core/component/idService"
	"eassy/core/component/msgService"
	"eassy/core/component/pkgService"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

/*****************客户端Manager*********************/
type CliMgr struct {
	CliMap map[int64]*Cli
	Lock   sync.RWMutex
}

type ICliMgr interface {
	New(ws *websocket.Conn) *Cli
	Destroy(cliId int64)
	GetCliIdByWs(ws *websocket.Conn) int64
}

func (p *CliMgr) New(ws *websocket.Conn) *Cli {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	cliId := idService.GenerateID().Int64()
	p.CliMap[cliId] = &Cli{
		Id:         cliId,
		Conn:       ws,
		ReqMsgChan: make(chan *MsgInfo, 100),
		AliveChan:  make(chan struct{}, 1),
	}
	return p.CliMap[cliId]
}

func (p *CliMgr) Destroy(cliId int64) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	delete(p.CliMap, cliId)
}

func (p *CliMgr) GetCliIdByWs(ws *websocket.Conn) int64 {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	for k, v := range p.CliMap {
		if v.Conn == ws {
			return k
		}
	}
	return 0
}

/*****************客户端*********************/
type Cli struct {
	Id             int64
	Conn           *websocket.Conn
	ServerId       uint
	ReqMsgChan     chan *MsgInfo
	AliveChan      chan struct{} //连接是否正常
	HeartBeatTimes int
	LastBeatTime   time.Time
}

type MsgInfo struct {
	Route int
	Body  []byte
}

type ICli interface {
	RecvData()
	HandleData()
	Dispatch(content []byte)
	HeartBeat()
	Send(protoId int, buffer []byte)
	Disconnect()
	Reconnect()
	Ticker() //心跳包监测  清理未连接cli
}

func (p *Cli) RecvData() {
	defer p.Disconnect()
	go p.Ticker()
	var (
		//msgType int
		content []byte
		err     error
		info    = new(MsgInfo)
	)
	for {
		_, content, err = p.Conn.ReadMessage()
		if err != nil {
			break
		}
		length := len(content)
		if length == 0 || length >= 4096 {
			break
		}
		info.Route, info.Body = pkgService.PkgDecode(content)
		//bodyLen := pkgService.GetPkgBodyLen(content)
		//if bodyLen+pkgService.PkgHeadBytes == length {
		p.ReqMsgChan <- info
		//}

	}
	return
}

func (p *Cli) HandleData() {
	for {
		select {
		case info, ok := <-p.ReqMsgChan:
			//fmt.Println(string(content))
			if !ok {
				//p.Send(websocket.CloseMessage,[]byte{})
				return
			}
			if info.Route == 0 {
				go p.HeartBeat()
			} else {
				go p.Dispatch(info.Route, info.Body)
			}
		}
	}
}

func (p *Cli) HeartBeat() {
	p.HeartBeatTimes = 0
	p.LastBeatTime = time.Now()
	p.Send([]byte("PONG"))
}

func (p *Cli) Reconnect() {
	//todo
}

func (p *Cli) Dispatch(route int, body []byte) {
	//todo: dispatch msg
	msgInfo, ok := msgService.GetMsgService().GetMsgByRouteId(route)
	if !ok {
		log.Print("未注册的route ！")
		return
	}
	//msgInfo.
	_ = msgInfo
	p.Send([]byte("hahaha"))
}

func (p *Cli) Send(buffer []byte) {
	err := p.Conn.WriteMessage(websocket.BinaryMessage, buffer)
	if err != nil {
		log.Fatal(err)
	}
	//p.Conn.Write(bytes)
}

func (p *Cli) Disconnect() {
	defer CliManager.Destroy(p.Id)
	err := p.Conn.Close()
	if err != nil {
		log.Print(err.Error())
	}
	p.AliveChan <- struct{}{}
	close(p.AliveChan)
	close(p.ReqMsgChan)
}

func (p *Cli) Ticker() {
	defer p.Conn.Close()
	var ti = 5 * time.Second
	var t = time.NewTicker(ti)
	for {
		select {
		case <-t.C:
			now := time.Now()
			if now.Sub(p.LastBeatTime) > ti {
				p.HeartBeatTimes++
			}
			if p.HeartBeatTimes >= 3 {
				return
			}
		case <-p.AliveChan:
			return
		}
	}
}
