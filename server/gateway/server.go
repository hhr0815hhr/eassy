package gateway

import (
	"eassy/core/server"
	"github.com/gin-gonic/gin"
)

func Run(s *server.Server) (err error) {
	r := getRouter()
	return s.Run(r)
}

func getRouter() *gin.Engine {
	r := gin.New()
	//r.Use(log.WebsocketLogger())
	setRouter(r)
	return r
}

func setRouter(r *gin.Engine) {
	r.GET("/eassy", websocketHandler)
}
