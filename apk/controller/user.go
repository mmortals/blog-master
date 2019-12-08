package controller

import (
	"blog-master/apk/service"
	"blog-master/public"
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
func (UserController) login(c *public.MyfContext) {

}

func (UserController) Delete(c *public.MyfContext) {
	service.UserService.Delete(c)
}
