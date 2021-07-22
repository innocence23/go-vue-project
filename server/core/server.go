package core

import (
	"fmt"
	"project/global"
	"project/handler"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	global.GVA_VP = initViper() // 初始化Viper
	global.GVA_LOG = initZap()  // 初始化zap日志库
	global.GVA_DB = initDB()    // gorm连接数据库
	global.GVA_REDIS = Redis()  // 初始化redis

	Router := handler.InitRouter()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
