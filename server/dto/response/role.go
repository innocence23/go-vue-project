package response

import "project/model/system"

type RoleResponse struct {
	Authority system.Role `json:"authority"`
}

type RoleCopyResponse struct {
	Authority      system.Role `json:"authority"`
	OldAuthorityId string      `json:"oldAuthorityId"` // 旧角色ID
}
