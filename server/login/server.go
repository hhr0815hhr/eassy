package login

import (
	"eassy/core/log"
	"eassy/core/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(s *server.Server) (err error) {
	r := getRouter()
	return s.Run(r)
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(log.HttpLogger(), AccessControlAllowOrigin())
	setRouter(r)
	return r
}

func AccessControlAllowOrigin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("origin"))
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
