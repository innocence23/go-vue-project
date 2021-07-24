package response

import (
	"project/entity"
	"project/model/system"
)

type SysMenusResponse struct {
	Menus []entity.Menu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu entity.Menu `json:"menu"`
}
