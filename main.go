package main

import (
	"fmt"
	"log"
	"syscall"
	"template/model"
	"template/pkg/logging"
	"template/pkg/setting"
	"template/router"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	model.Setup()
	logging.Setup()
}
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	// routerInit := router.InitRouter()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	// server := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
	// 	Handler:        routerInit,
	// 	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// server.ListenAndServe()

	server := endless.NewServer(endPoint, router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}
	logging.Info("运行成功...")
}
