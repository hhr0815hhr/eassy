package login

import "github.com/gin-gonic/gin"

func loginHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": nil,
		"msg":  "1111",
	})
}

func registerHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": nil,
		"msg":  "1111",
	})
}
