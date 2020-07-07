package etcd

import (
	"eassy/util"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type IRegistry interface {
	RegistryNode(node string) error
	UnRegistry()
}

//服务注册
func RegisterService(nodeType string, port string) {
	ok, serverIP := util.ServerIP()
	if !ok {
		serverIP = "127.0.0.1"
	}
	serviceInfo := ServiceInfo{
		Name: nodeType,
		Addr: serverIP + ":" + port,
	}
	regService, err := newService(serviceInfo, etcdSlice)
	if err != nil {
		log.Fatal(err)
	}
	go regService.Start()
}

func newService(info ServiceInfo, endpoints []string) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Service{
		Info:   info,
		stop:   make(chan error),
		client: cli,
	}, err
}
