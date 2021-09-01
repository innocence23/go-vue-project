package entity

import (
	"project/zvar"
)

type User struct {
	zvar.Model
	Username    string   `json:"userName" gorm:"comment:用户登录名"`                   // 用户登录名
	Password    string   `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName    string   `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	Avatar      string   `json:"avatar" gorm:"comment:用户头像"`                      // 用户头像
	RoleIds     []string `json:"roleIds" gorm:"-"`                                // 用户角色ID
	SideMode    string   `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`     // 用户侧边主题
	ActiveColor string   `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"` // 活跃颜色
	BaseColor   string   `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`      // 基础颜色
}
