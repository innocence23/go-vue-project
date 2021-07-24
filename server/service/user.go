package service

import (
	"errors"
	"project/dto/request"
	"project/model/system"
	"project/utils"
	"project/zvar"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

func (userService *UserService) Register(u system.User) (err error, userInter system.User) {
	var user system.User
	if !errors.Is(zvar.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	u.Password, _ = utils.HashPassword(u.Password)
	u.UUID = uuid.NewV4()
	err = zvar.DB.Create(&u).Error
	return err, u
}

func (userService *UserService) Login(u *system.User) (err error, userInter *system.User) {
	var user system.User
	u.Password = utils.Md5([]byte(u.Password))
	err = zvar.DB.Where("username = ?", u.Username).Preload("Authority").First(&user).Error
	if err != nil {
		return err, &user
	}
	match, err := utils.ComparePasswords(user.Password, u.Password)
	if err != nil {
		return err, &user
	}
	if !match {
		return errors.New("用户名密码不正确已注册"), &user
	}
	return err, &user
}

func (userService *UserService) ChangePassword(u *system.User, newPassword string) (err error, userInter *system.User) {
	var user system.User
	u.Password, _ = utils.HashPassword(u.Password)
	newPassword, _ = utils.HashPassword(newPassword)
	err = zvar.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", newPassword).Error
	return err, u
}

func (userService *UserService) List(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&system.User{})
	var userList []system.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) SetRole(uuid uuid.UUID, roleId string) (err error) {
	err = zvar.DB.Where("uuid = ?", uuid).First(&system.User{}).Update("authority_id", roleId).Error
	return err
}

func (userService *UserService) Delete(id float64) (err error) {
	var user system.User
	err = zvar.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

func (userService *UserService) Update(reqUser system.User) (err error, user system.User) {
	err = zvar.DB.Updates(&reqUser).Error
	return err, reqUser
}

func (userService *UserService) FindById(id int) (err error, user *system.User) {
	var u system.User
	err = zvar.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (userService *UserService) FindByUuid(uuid string) (err error, user *system.User) {
	var u system.User
	if err = zvar.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
