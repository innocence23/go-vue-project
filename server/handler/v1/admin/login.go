package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/entity"
	"project/handler/middleware"
	"project/model/system"
	"project/service"
	"project/utils"
	"project/zvar"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type baseHandler struct {
	userService *service.UserService
	jwtService  *service.JwtService
	rbacService *service.RbacService
	roleService *service.RoleService
}

func NewBaseHandler() *baseHandler {
	return &baseHandler{
		userService: &service.UserService{},
		rbacService: &service.RbacService{},
		jwtService:  &service.JwtService{},
		roleService: &service.RoleService{},
	}
}

func (h *baseHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("basic").Use(middleware.OperationRecord())
	{
		apiRouter.POST("login", h.login)
		apiRouter.POST("captcha", h.captcha)
	}
}

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (h *baseHandler) captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(zvar.Config.Captcha.ImgHeight, zvar.Config.Captcha.ImgWidth, zvar.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		zvar.Log.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (h *baseHandler) login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaId, req.Captcha, true) {
		u := &system.User{Username: req.Username, Password: req.Password}
		if err, user := h.userService.Login(u); err != nil {
			zvar.Log.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			roleIds, _ := h.rbacService.GetRolesForUser(user.Username)
			user.Role, _ = h.roleService.FindByIds(roleIds)
			h.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (h *baseHandler) tokenNext(c *gin.Context, user system.User) {
	j := &middleware.JWT{SigningKey: []byte(zvar.Config.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		NickName:   user.NickName,
		Username:   user.Username,
		RoleId:     user.RoleId,
		BufferTime: zvar.Config.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                        // 签名生效时间
			ExpiresAt: time.Now().Unix() + zvar.Config.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                        // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zvar.Log.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !zvar.Config.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := h.jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := h.jwtService.SetRedisJWT(token, user.Username); err != nil {
			zvar.Log.Error("设置登录状态失败!", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		zvar.Log.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT entity.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := h.jwtService.InBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := h.jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
