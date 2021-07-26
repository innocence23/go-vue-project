package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID         int
	Username   string
	NickName   string
	RoleId     []string
	BufferTime int64
	jwt.StandardClaims
}
