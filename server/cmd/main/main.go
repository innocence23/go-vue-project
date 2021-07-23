package main

import (
	"fmt"
	"project/handler"
	"project/pkg"
	"project/zvar"
	"time"

	"github.com/fvbock/endless"
	"go.uber.org/zap"
)

func main() {
	configFile := "./config.yaml"

	zvar.Viper = pkg.InitViper(configFile) // 初始化Viper
	zvar.Log = pkg.InitZap()               // 初始化zap日志库
	zvar.DB = pkg.InitDB()                 // gorm连接数据库
	zvar.Redis = pkg.InitRedis()           // 初始化redis

	router := handler.InitRouter()
	router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", zvar.Config.System.Addr)

	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	time.Sleep(10 * time.Microsecond)
	zvar.Log.Info("server run success on ", zap.String("address", address))
	zvar.Log.Error(s.ListenAndServe().Error())
}
