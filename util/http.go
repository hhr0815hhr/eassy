package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	respCodeSuccess = iota
	respCodeFailed
)

func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": respCodeSuccess,
		"msg":  msg,
		"data": data,
	})
}

func Failed(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": respCodeFailed,
		"msg":  msg,
		"data": data,
	})
}
