package entity

import "project/zvar"

type Role struct {
	zvar.Model
	AuthorityId     string `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleId          string `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`      // 角色ID
	AuthorityName   string `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	RoleName        string `json:"roleName" gorm:"comment:角色名"`                                         // 角色名
	ParentId        string `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []Role `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []Role `json:"children" gorm:"-"`
	Menus           []Menu `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}
