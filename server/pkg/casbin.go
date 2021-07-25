package pkg

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("config/casbin.conf", adapter) //todo配置
	if err != nil {
		panic(err)
	}
	// 日志记录
	//enforcer.EnableLog(true)
	return enforcer
}
