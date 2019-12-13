package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 框架实例
type MyfApp struct {
	Gin *gin.Engine
}

func New() (app *MyfApp, err error) {

	// 创建APP
	app = &MyfApp{}

	// 创建Gin
	app.initGin()
	// 创建客户端

	// 注册中间件
	return
}

// 初始化Gin
func (app *MyfApp) initGin() {

	// 调试模式
	//ddd

	// 生产模式
	app.Gin = gin.New()
}

// 启动APP
func (app *MyfApp) Run() (err error) {

	srv := &http.Server{
		Addr:         "localhost:8081",
		Handler:      app.Gin,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//启动server服务
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errMsg := fmt.Sprintf("启动server失败：%+v, %+v ", srv, err)
			panic(errMsg)
		}
	}()

	//等待优雅退出命令
	app.waitGraceExit(srv)
	return

}

// 等待优雅退出
func (app *MyfApp) waitGraceExit(server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Fprintf(os.Stdout, "收到信号: %s, 服务正在退出... \n", s.String())
			server.Close()
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
