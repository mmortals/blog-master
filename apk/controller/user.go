package controller

import (
	"blog-master/apk/clients/myredis"
	"blog-master/apk/service"
	"blog-master/public"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始化 user 路由组
func init() {
	//g := Engine.Group("/user")
	//g.POST("/register", public.Handler(UserController{}.Register))
	//g.POST("/delete/:id", public.Handler(UserController{}.Delete))

}

type UserController struct{}

// Register 用户注册
func (UserController) Register(c *public.MyfContext) {
	service.UserService.Register(c)
}

//用户登录
func (UserController) Login(c *public.MyfContext) {

	user, err := service.UserService.Login(c)

	if nil != err {
		fmt.Println(err)
		c.Gin.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "userController.Login  error ",
			"error":   err,
		})
	} else {
		myredis.MyfRedis.Set("userId", user.Id, 0)
		c.Gin.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": user,
			"error":   err,
			"success": "success",
		})

	}

}

func (UserController) Delete(c *public.MyfContext) {
	service.UserService.Delete(c)
}

func (UserController) FindUserById(c *public.MyfContext) {
	user, err := service.UserService.FindUserById(c)

	if nil != err {
		fmt.Println(err)
		c.Gin.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "userController.Login  error ",
			"error":   err,
		})
	} else {
		myredis.MyfRedis.Set("userId", user.Id, 0)
		c.Gin.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": user,
			"error":   err,
			"success": "success",
		})
	}
}
