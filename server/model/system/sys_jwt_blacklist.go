package system

import "project/zvar"

type JwtBlacklist struct {
	zvar.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
