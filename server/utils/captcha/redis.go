package captcha

import (
	"context"
	"project/zvar"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() base64Captcha.Store {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) {
	err := zvar.Redis.Set(context.Background(), rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		zvar.Log.Error("RedisStoreSetError!", zap.Error(err))
	}
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := zvar.Redis.Get(context.Background(), key).Result()
	if err != nil {
		zvar.Log.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := zvar.Redis.Del(context.Background(), key).Err()
		if err != nil {
			zvar.Log.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
