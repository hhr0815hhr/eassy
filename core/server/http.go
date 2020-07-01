package server

import (
	"eassy/core/etcd"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func httpRun(s *Server, router *gin.Engine) error {
	//etcd服务注册
	etcd.RegisterService(s.Type, s.Port)
	ginpprof.Wrap(router)
	err := router.Run(":" + s.Port)
	return err
}
