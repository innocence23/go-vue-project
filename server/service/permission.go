package service

import (
	"errors"
	"project/dto/request"
	"project/entity"
	"project/zvar"

	"gorm.io/gorm"
)

type PermissionService struct {
}

func (permissionService *PermissionService) Create(perm entity.Permission) (err error) {
	if !errors.Is(zvar.DB.Where("path = ? AND method = ?", perm.Path, perm.Method).First(&entity.Permission{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同perm")
	}
	return zvar.DB.Create(&perm).Error
}

func (permissionService *PermissionService) Delete(perm entity.Permission) (err error) {
	err = zvar.DB.Delete(&perm).Error
	return err
}

func (permissionService *PermissionService) List(perm entity.Permission, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&entity.Permission{})
	var permList []entity.Permission

	if perm.Path != "" {
		db = db.Where("path LIKE ?", "%"+perm.Path+"%")
	}
	if perm.Description != "" {
		db = db.Where("description LIKE ?", "%"+perm.Description+"%")
	}
	if perm.Method != "" {
		db = db.Where("method = ?", perm.Method)
	}
	if perm.Group != "" {
		db = db.Where("`group` = ?", perm.Group)
	}
	err = db.Count(&total).Error

	if err != nil {
		return
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&permList).Error
		} else {
			err = db.Order("`group`").Find(&permList).Error
		}
	}
	return
}

func (permissionService *PermissionService) ListNoLimit() (perms []entity.Permission, err error) {
	err = zvar.DB.Limit(1000).Find(&perms).Error //限制最多取一千条
	return
}

func (permissionService *PermissionService) Show(id int) (perm entity.Permission, err error) {
	err = zvar.DB.First(&perm, id).Error
	return
}
