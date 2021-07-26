package service

import (
	"errors"
	"project/dto/request"
	"project/entity"
	"project/utils"
	"project/zvar"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserService struct {
}

func (userService *UserService) Register(u entity.User) (user entity.User, err error) {
	if !errors.Is(zvar.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		err = errors.New("用户名已注册")
		return
	}
	u.Password, _ = utils.HashPassword(u.Password)
	err = zvar.DB.Create(&u).Error
	user = u
	return
}

func (userService *UserService) Login(u *entity.User) (user *entity.User, err error) {
	err = zvar.DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return
	}
	match, err := utils.ComparePasswords(user.Password, u.Password)
	if err != nil {
		return
	}
	if !match {
		err = errors.New("用户名密码不正确")
		return
	}
	return
}

func (userService *UserService) ChangePassword(u *entity.User, newPassword string) (user *entity.User, err error) {
	originUser, err := userService.Login(u)
	if err != nil {
		return
	}
	newPassword, _ = utils.HashPassword(newPassword)
	err = zvar.DB.Where("id = ?", originUser.ID).First(&user).Update("password", newPassword).Error
	return
}

func (userService *UserService) List(info request.PageInfo) (userList []entity.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&entity.User{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	if err != nil {
		return
	}
	var newList []entity.User
	for _, v := range userList {
		roleIds, _ := (&RbacService{}).GetRolesForUser(cast.ToString(v.ID))
		v.Roles, err = (&RoleService{}).FindByIds(cast.ToIntSlice(roleIds))
		newList = append(newList, v)
	}
	userList = newList
	return
}

func (userService *UserService) Delete(id int) (err error) {
	var user entity.User
	err = zvar.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

func (userService *UserService) Update(reqUser entity.User) (user entity.User, err error) {
	err = zvar.DB.Updates(&reqUser).Error
	user = reqUser
	return
}

func (userService *UserService) FindById(id int) (user *entity.User, err error) {
	err = zvar.DB.First(&user, id).Error
	return
}
