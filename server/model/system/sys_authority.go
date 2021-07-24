package system

import (
	"time"
)

type Role struct {
	CreatedAt     time.Time     // 创建时间
	UpdatedAt     time.Time     // 更新时间
	DeletedAt     *time.Time    `sql:"index"`
	RoleId        string        `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleName      string        `json:"roleName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      string        `json:"parentId" gorm:"comment:父角色ID"`                                  // 父角色ID
	DataRoleId    []Role        `json:"dataRoleId" gorm:"many2many:sys_data_authority_id"`
	Children      []Role        `json:"children" gorm:"-"`
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter string        `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}
