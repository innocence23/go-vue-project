package admin

import (
	"fmt"
	"project/dto/request"
	"project/dto/response"
	"project/model/system"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type userHandler struct {
	userService *service.UserService
	rbacService *service.RbacService
	roleService *service.RoleService
}

func NewUserHandler() *userHandler {
	return &userHandler{
		userService: &service.UserService{},
		rbacService: &service.RbacService{},
		roleService: &service.RoleService{},
	}
}

func (h *userHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("user")
	{
		apiRouter.POST("register", h.register)
		apiRouter.POST("changePassword", h.changePassword)
		apiRouter.POST("list", h.list)
		apiRouter.DELETE("delete", h.delete)
		apiRouter.PUT("update", h.update)
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/user/register"] = zvar.RouteInfo{Group: "user", Name: "注册账号"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/user/changePassword"] = zvar.RouteInfo{Group: "user", Name: "修改密码"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/user/list"] = zvar.RouteInfo{Group: "user", Name: "用户列表"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/user/delete"] = zvar.RouteInfo{Group: "user", Name: "删除用户"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/user/update"] = zvar.RouteInfo{Group: "user", Name: "更新用户"}
}

// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (h *userHandler) register(c *gin.Context) {
	var req request.Register
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.User{Username: req.Username, NickName: req.NickName, Password: req.Password, HeaderImg: req.HeaderImg, RoleId: req.RoleId}
	err, userReturn := h.userService.Register(*user)
	if err != nil {
		zvar.Log.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(response.UserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.UserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (h *userHandler) changePassword(c *gin.Context) {
	var req request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.User{Username: req.Username, Password: req.Password}
	if err, _ := h.userService.ChangePassword(u, req.NewPassword); err != nil {
		zvar.Log.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/list [post]
func (h *userHandler) list(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := h.userService.List(pageInfo); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		var newList []system.User
		for _, v := range list {
			roleIds, _ := h.rbacService.GetRolesForUser(cast.ToString(v.ID))
			v.Roles, err = h.roleService.FindByIds(roleIds)
			fmt.Println("--------", roleIds, v.Roles, err)
			newList = append(newList, v)
		}

		response.OkWithDetailed(response.PageResult{
			List:     newList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/delete [delete]
func (h *userHandler) delete(c *gin.Context) {
	var req request.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(req.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := h.userService.Delete(req.ID); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags User
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.User true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/update [put]
func (h *userHandler) update(c *gin.Context) {
	var req system.User
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := h.userService.Update(req); err != nil {
		zvar.Log.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}
