package service

import (
	"errors"
	"project/dto/request"
	"project/dto/response"
	"project/model/system"
	"project/zvar"
	"strconv"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.Role
//@return: err error, authority model.Role

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.Role) (err error, authority system.Role) {
	var authorityBox system.Role
	if !errors.Is(zvar.DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = zvar.DB.Create(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.RoleCopyResponse
//@return: err error, authority model.Role

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.RoleCopyResponse) (err error, authority system.Role) {
	var authorityBox system.Role
	if !errors.Is(zvar.DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []system.Role{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = zvar.DB.Create(&copyInfo.Authority).Error

	paths := (&CasbinService{}).GetPermByRoleId(copyInfo.OldAuthorityId)
	err = (&CasbinService{}).Update(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.Role
//@return: err error, authority model.Role

func (authorityService *AuthorityService) UpdateAuthority(auth system.Role) (err error, authority system.Role) {
	err = zvar.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.Role{}).Updates(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.Role
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.Role) (err error) {
	if !errors.Is(zvar.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(zvar.DB.Where("parent_id = ?", auth.AuthorityId).First(&system.Role{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := zvar.DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = zvar.DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	(&CasbinService{}).ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB
	var authority []system.Role
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.Role
//@return: err error, sa model.Role

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.Role) (err error, sa system.Role) {
	err = zvar.DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.Role
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(auth system.Role) error {
	var s system.Role
	zvar.DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := zvar.DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.Role
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.Role) error {
	var s system.Role
	zvar.DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := zvar.DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.Role
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.Role) (err error) {
	err = zvar.DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
