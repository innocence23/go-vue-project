package service

import (
	"project/dto/request"
	"project/model/system"
	"project/zvar"
)

type RoleService struct {
}

func (roleService *RoleService) FindByIds(ids []string) (roleList []system.Role, err error) {
	err = zvar.DB.Where("authority_id in ?", ids).Find(&roleList).Error
	return roleList, err
}

func (roleService *RoleService) List(info request.PageInfo) (list []system.Role, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB
	var roleList []system.Role
	err = db.Limit(limit).Offset(offset).Where("parent_id = 0").Find(&roleList).Error
	return roleList, total, err
}
