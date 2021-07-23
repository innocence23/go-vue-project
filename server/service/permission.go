package service

import (
	"errors"
	"project/dto/request"
	"project/model/system"
	"project/zvar"

	"gorm.io/gorm"
)

type PermissionService struct {
}

func (permissionService *PermissionService) Create(perm system.Permission) (err error) {
	if !errors.Is(zvar.DB.Where("path = ? AND method = ?", perm.Path, perm.Method).First(&system.Permission{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同perm")
	}
	return zvar.DB.Create(&perm).Error
}

func (permissionService *PermissionService) Delete(perm system.Permission) (err error) {
	err = zvar.DB.Delete(&perm).Error
	CasbinServiceApp.ClearCasbin(1, perm.Path, perm.Method)
	return err
}

func (permissionService *PermissionService) List(perm system.Permission, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&system.Permission{})
	var permList []system.Permission

	if perm.Path != "" {
		db = db.Where("path LIKE ?", "%"+perm.Path+"%")
	}

	if perm.Description != "" {
		db = db.Where("description LIKE ?", "%"+perm.Description+"%")
	}

	if perm.Method != "" {
		db = db.Where("method = ?", perm.Method)
	}

	if perm.ApiGroup != "" {
		db = db.Where("api_group = ?", perm.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, permList, total
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
			err = db.Order("api_group").Find(&permList).Error
		}
	}
	return err, permList, total
}

func (permissionService *PermissionService) ListNoLimit() (err error, perms []system.Permission) {
	err = zvar.DB.Find(&perms).Error
	return
}

func (permissionService *PermissionService) Show(id float64) (err error, perm system.Permission) {
	err = zvar.DB.Where("id = ?", id).First(&perm).Error
	return
}

func (permissionService *PermissionService) Update(perm system.Permission) (err error) {
	var oldA system.Permission
	err = zvar.DB.Where("id = ?", perm.ID).First(&oldA).Error
	if oldA.Path != perm.Path || oldA.Method != perm.Method {
		if !errors.Is(zvar.DB.Where("path = ? AND method = ?", perm.Path, perm.Method).First(&system.Permission{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同perm路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, perm.Path, oldA.Method, perm.Method)
		if err != nil {
			return err
		} else {
			err = zvar.DB.Save(&perm).Error
		}
	}
	return err
}

func (permissionService *PermissionService) DeleteByIds(ids []int) (err error) {
	return zvar.DB.Delete(&system.Permission{}, ids).Error
}
