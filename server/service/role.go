package service

import (
	"project/model/system"
	"project/zvar"
)

type RoleService struct {
}

func (roleService *RoleService) FindByIds(ids []string) (roleList []system.Role, err error) {
	err = zvar.DB.Where("authority_id in ?", ids).Find(&roleList).Error
	return roleList, err
}
