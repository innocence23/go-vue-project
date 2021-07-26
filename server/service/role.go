package service

import (
	"project/dto/request"
	"project/entity"
	"project/zvar"
)

type RoleService struct {
}

func (roleService *RoleService) FindByIds(ids []int) (roleList []entity.Role, err error) {
	err = zvar.DB.Find(&roleList, ids).Error
	return roleList, err
}

func (roleService *RoleService) List(info request.PageInfo) (roleList []entity.Role, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&entity.Role{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Where("parent_id = 0").Find(&roleList).Error
	return
}

func (roleService *RoleService) Create(role entity.Role) (err error) {
	err = zvar.DB.Create(&role).Error
	return
}

func (roleService *RoleService) Update(role entity.Role) (err error) { //todo 改成map更新
	err = zvar.DB.Updates(role).Error
	return
}

func (roleService *RoleService) Delete(id int) (err error) {
	err = zvar.DB.Delete(&entity.Role{}, id).Error
	return
}
