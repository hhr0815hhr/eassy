package login

import "github.com/gin-gonic/gin"

func setRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/login", loginHandler)
		api.POST("/register", registerHandler)
	}
}
