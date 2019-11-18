package controller

import (
	"blog-master/apk/service"
	"blog-master/public"
)

//初始化 user 路由组
func init() {
	g := Engine.Group("/user")
	g.GET("/regist/:id/:username", public.Handler(UserController{}.Regist))
}

type UserController struct{}

// Regist 用户注册
func (UserController) Regist(c *public.MyfContext) {
	service.UserService.Regist(c)
}
