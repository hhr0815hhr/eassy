package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Type     string
	Port     string
	Protocol string //http websocket socket
}

type InsServer interface {
	Run(router *gin.Engine) error
	Shutdown(ctx context.Context) (err error)
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	return
}

func (s *Server) Run(router *gin.Engine) (err error) {
	switch s.Protocol {
	case "http", "websocket":
		err = httpRun(s, router)
	case "socket":
		err = socketRun(s)
	}
	return
}
