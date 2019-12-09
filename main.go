package main

import (
	"blog-master/apk/app"
	"blog-master/apk/controller"
	"blog-master/public"
	"fmt"
)

func main() {

	fmt.Println(">>>>>go go go !<<<<<")
	// 启动web容器
	//_ = controller.Engine.Run()

	app, _ := app.New()
	//test
	

	g := app.Gin.Group("/user")
	g.POST("/register", public.Handler(controller.UserController{}.Register))
	g.POST("/delete/:id", public.Handler(controller.UserController{}.Delete))
	app.Run()
}
