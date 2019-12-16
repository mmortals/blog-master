package public

import (
	"blog-master/apk/clients/myredis"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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

		if c.Request.RequestURI != "/user/login" {

			userId, err := myredis.MyfRedis.Get("userId").Result()

			if "" == userId || nil != err {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			}

		}
		// 请求超时控制
		timeoutCtx, cancelFunc := context.WithTimeout(c, time.Duration(5000)*time.Millisecond)
		defer cancelFunc()
		context := new(MyfContext)
		context.Gin = c
		context.Context = timeoutCtx
		handler(context)
	}
}
