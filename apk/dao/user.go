package dao

import (
	"blog-master/public"
	"fmt"
	"github.com/gin-gonic/gin"
)

type userDao struct{}

var UserDao = new(userDao)

func (*userDao) Regist(c *public.MyfContext) {

	var url = c.Gin.Request.RequestURI

	fmt.Printf(">>>>>url is :<<<<<", url)

	c.Gin.JSON(200, gin.H{
		"message": url,
	})
}
