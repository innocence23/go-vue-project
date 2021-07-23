package pkg

import (
	"context"
	"project/zvar"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() *redis.Client {
	redisCfg := zvar.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zvar.Log.Error("redis connect ping failed, err:", zap.Any("err", err))
	}
	return client
}
