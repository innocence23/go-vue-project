package core

import (
	"context"
	"project/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	}
	return client
}
