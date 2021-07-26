package service

import (
	"context"
	"errors"
	"project/entity"
	"project/zvar"
	"time"

	"gorm.io/gorm"
)

type JwtService struct {
}

func (jwtService *JwtService) InBlacklist(jwtList entity.JwtBlacklist) (err error) {
	err = zvar.DB.Create(&jwtList).Error
	return
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := zvar.DB.Where("jwt = ?", jwt).First(&entity.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = zvar.Redis.Get(context.Background(), userName).Result()
	return
}

func (jwtService *JwtService) SetRedisJWT(jwt, userName string) (err error) {
	timer := time.Duration(zvar.Config.JWT.ExpiresTime) * time.Second
	err = zvar.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return
}
