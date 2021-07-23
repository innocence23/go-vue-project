package zvar

import (
	"project/config"
	"project/utils/timer"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	Redis              *redis.Client
	Config             config.Server
	Viper              *viper.Viper
	Log                *zap.Logger
	Timer              timer.Timer = timer.NewTimerTask()
	ConcurrencyControl             = &singleflight.Group{}
)
