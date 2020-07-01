package gateway

import (
	"eassy/core/component/idService"
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
	Destroy(ws *websocket.Conn)
	GetCliIdByWs(ws *websocket.Conn) int64
}

func (p *CliMgr) New(ws *websocket.Conn) *Cli {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	cliId := idService.GenerateID().Int64()
	p.CliMap[cliId] = &Cli{
		Id:        cliId,
		Conn:      ws,
		MsgChan:   make(chan []byte, 100),
		AliveChan: make(chan struct{}, 1),
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
	MsgChan        chan []byte
	AliveChan      chan struct{} //连接是否正常
	HeartBeatTimes int
	LastBeatTime   time.Time
}

type ICli interface {
	RecvData()
	HandleData()
	Dispatch()
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
	)
	for {
		_, content, err = p.Conn.ReadMessage()
		if err != nil {
			break
		}
		if len(content) == 0 || len(content) >= 4096 {
			break
		}
		p.MsgChan <- content
	}
	return
}

func (p *Cli) HandleData() {
	for {
		select {
		case content, ok := <-p.MsgChan:
			//fmt.Println(string(content))
			if !ok {
				//p.Send(websocket.CloseMessage,[]byte{})
				return
			}
			if string(content) == "PING" {
				go p.HeartBeat()
			} else {
				go p.Dispatch()
			}
		}
	}
}

func (p *Cli) HeartBeat() {
	p.HeartBeatTimes = 0
	p.LastBeatTime = time.Now()
	p.Send(websocket.PongMessage, []byte("PONG"))
}

func (p *Cli) Reconnect() {
	//todo
}

func (p *Cli) Dispatch() {
	//todo: dispatch msg
	p.Send(websocket.TextMessage, []byte("hahaha"))
}

func (p *Cli) Send(protoId int, buffer []byte) {
	err := p.Conn.WriteMessage(protoId, buffer)
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
	close(p.MsgChan)
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
