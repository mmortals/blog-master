package public

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// Myf框架请求上下文
type MyfContext struct {
	context.Context
	Gin *gin.Context
}

// Myf框架controller定义
type MyfHandleFunc func(c *MyfContext)

func Handler(handler MyfHandleFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		// 请求超时控制
		timeoutCtx, cancelFunc := context.WithTimeout(c, time.Duration(5000)*time.Millisecond)
		defer cancelFunc()
		context := new(MyfContext)
		context.Gin = c
		context.Context = timeoutCtx
		handler(context)
	}
}
