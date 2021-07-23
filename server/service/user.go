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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: err error, userInter model.User

type UserService struct {
}

func (userService *UserService) Register(u system.User) (err error, userInter system.User) {
	var user system.User
	if !errors.Is(zvar.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = zvar.DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.User
//@return: err error, userInter *model.User

func (userService *UserService) Login(u *system.User) (err error, userInter *system.User) {
	var user system.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = zvar.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.User, newPassword string
//@return: err error, userInter *model.User

func (userService *UserService) ChangePassword(u *system.User, newPassword string) (err error, userInter *system.User) {
	var user system.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = zvar.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&system.User{})
	var userList []system.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = zvar.DB.Where("uuid = ?", uuid).First(&system.User{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.User
	err = zvar.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.User
//@return: err error, user model.User

func (userService *UserService) SetUserInfo(reqUser system.User) (err error, user system.User) {
	err = zvar.DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.User

func (userService *UserService) FindUserById(id int) (err error, user *system.User) {
	var u system.User
	err = zvar.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.User

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.User) {
	var u system.User
	if err = zvar.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
