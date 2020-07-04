package server

import (
	"eassy/core/server"
	"eassy/server/game"
	"eassy/server/gateway"
	"eassy/server/hall"
	"eassy/server/login"
	"errors"
	"github.com/spf13/viper"
	"sync"
)

var s *server.Server
var once sync.Once

func NewServer(serverType, serverPort string) *server.Server {
	//serverProtocol := viper.GetString(serverType+".protocol")
	serverProtocol := viper.GetString("server." + serverType + ".protocol")
	once.Do(func() {
		s = new(server.Server)
		s.Type = serverType
		s.Port = serverPort
		s.Protocol = serverProtocol
	})
	return s
}

func Run(server *server.Server) (err error) {
	switch server.Type {
	case "gate":
		err = gateway.Run(server)
	case "login":
		err = login.Run(server)
	case "game":
		err = game.Run(server)
	case "hall":
		err = hall.Run(server)
	default:
		err = errors.New("无效的服务类型！")
	}
	return
}
