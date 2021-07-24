package response

import "project/model/system"

type RoleResponse struct {
	Authority system.Role `json:"authority"`
}

type RoleCopyResponse struct {
	Authority system.Role `json:"authority"`
	OldRoleId string      `json:"oldRoleId"` // 旧角色ID
}
