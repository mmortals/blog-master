package main

import (
	"blog-master/apk/controller"
	"fmt"
)

func main() {

	fmt.Println(">>>>>go go go !<<<<<")
	// 启动web容器
	controller.Engine.Run()
}
