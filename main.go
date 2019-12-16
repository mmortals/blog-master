package main

import (
	"blog-master/apk/app"
	"blog-master/apk/controller"
	"blog-master/public"
	"fmt"
	"log"
)

func main() {

	fmt.Println(">>>>>go go go !<<<<<")
	// 启动web容器
	//_ = controller.Engine.Run()

	app, err := app.New()
	if nil != err {
		log.Fatal(">>>app start failed<<<")
	}

	g := app.Gin.Group("/user")
	g.POST("/register", public.Handler(controller.UserController{}.Register))
	g.POST("/delete/:id", public.Handler(controller.UserController{}.Delete))
	g.POST("/login", public.Handler(controller.UserController{}.Login))
	g.POST("/findUserById", public.Handler(controller.UserController{}.FindUserById))
	app.Run()
}
