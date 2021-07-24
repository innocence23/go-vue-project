package request

import (
	"project/model/system"
	"project/zvar"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus  []system.SysBaseMenu
	RoleId string // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		Model:     zvar.Model{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
