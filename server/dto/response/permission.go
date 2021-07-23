package response

import "project/model/system"

type PermissionResponse struct {
	Permission system.Permission `json:"permission"`
}

type PermissionListResponse struct {
	Permissions []system.Permission `json:"permissions"`
}
