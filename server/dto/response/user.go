package response

import (
	"project/entity"
)

type UserResponse struct {
	User entity.User `json:"user"`
}

type LoginResponse struct {
	User      entity.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
