package response

import (
	"project/entity"
)

type RoleResponse struct {
	Role entity.Role `json:"role"`
}
