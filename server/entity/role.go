package entity

import "project/zvar"

type Role struct {
	zvar.Model
	Name     string `json:"name" gorm:"comment:角色名"`       // 角色名
	ParentId string `json:"parentId" gorm:"comment:父角色ID"` // 父角色ID
	//DataAuthorityId []Role `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children []Role `json:"children" gorm:"-"`
	//Menus    []Menu `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
