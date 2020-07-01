package server

import (
	"eassy/core/etcd"
	"net"
)

func socketRun(s *Server) (err error) {
	listen, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic(err)
		return
	}
	//etcd服务注册
	etcd.RegisterService(s.Type, s.Port)
	//grpc
	_ = listen
	//registerService(&service.GameService{})

	return
}
