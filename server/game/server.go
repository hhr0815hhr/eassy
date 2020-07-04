package game

import (
	"eassy/core/etcd"
	g "eassy/core/grpc"
	"eassy/core/server"
	"net"
)

func Run(s *server.Server) (err error) {
	var listen net.Listener
	listen, err = net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic(err)
		return
	}

	//注册etcd服务
	etcd.RegisterService(s.Type, s.Port)

	//grpc服务
	err = g.GRPC.Serve(listen)
	return
	return
}
