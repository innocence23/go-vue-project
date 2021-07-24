package response

import (
	"project/entity"
)

type PermissionResponse struct {
	Permission entity.Permission `json:"permission"`
}

type PermissionListResponse struct {
	Permissions []entity.Permission `json:"permissions"`
}
