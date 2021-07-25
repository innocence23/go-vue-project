package admin

import (
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

type rolexxxHandler struct {
	service     *service.AuthorityService
	serviceMenu *service.MenuService
	roleService *service.RoleService
	rbacService *service.RbacService

	//serviceCasbin *service.CasbinService
}

func NewRolexxxHandler() *rolexxxHandler {
	return &rolexxxHandler{
		service:     &service.AuthorityService{},
		serviceMenu: &service.MenuService{},
		roleService: &service.RoleService{},
		rbacService: &service.RbacService{},
		//serviceCasbin: &service.CasbinService{},
	}
}

func (h *rolexxxHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("role")
	{
		apiRouter.POST("list", h.list)               // 获取角色列表
		apiRouter.POST("setUserRole", h.setUserRole) // 设置用户角色

		apiRouter.POST("createAuthority", h.CreateAuthority)   // 创建角色
		apiRouter.POST("deleteAuthority", h.DeleteAuthority)   // 删除角色
		apiRouter.PUT("updateAuthority", h.UpdateAuthority)    // 更新角色
		apiRouter.POST("copyAuthority", h.CopyAuthority)       // 更新角色
		apiRouter.POST("setDataAuthority", h.SetDataAuthority) // 设置角色资源权限
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/list"] = zvar.RouteInfo{Group: "role", Name: "角色列表"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/setUserRole"] = zvar.RouteInfo{Group: "role", Name: "设置用户角色"}

}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/list [post]
func (h *rolexxxHandler) list(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := h.roleService.List(pageInfo); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserRole true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setRole [post]
func (h *rolexxxHandler) setUserRole(c *gin.Context) {
	var req request.SetUserRole
	_ = c.ShouldBindJSON(&req)
	if UserVerifyErr := utils.Verify(req, utils.SetUserRoleorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if _, err := h.rbacService.AddRoleForUser(cast.ToString(req.ID), req.RoleId); err != nil {
		zvar.Log.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (h *rolexxxHandler) CreateAuthority(c *gin.Context) {
	var req system.Role
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := h.service.CreateAuthority(req); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = h.serviceMenu.AddMenuAuthority(request.DefaultMenu(), req.AuthorityId)
		//_ = h.serviceCasbin.Update(req.RoleId, request.DefaultCasbin())
		response.OkWithDetailed(response.RoleResponse{Authority: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.RoleCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (h *rolexxxHandler) CopyAuthority(c *gin.Context) {
	var copyInfo response.RoleCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := h.service.CopyAuthority(copyInfo); err != nil {
		zvar.Log.Error("拷贝失败!", zap.Any("err", err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.RoleResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (h *rolexxxHandler) DeleteAuthority(c *gin.Context) {
	var authority system.Role
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (h *rolexxxHandler) UpdateAuthority(c *gin.Context) {
	var auth system.Role
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authority := h.service.UpdateAuthority(auth); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.RoleResponse{Authority: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (h *rolexxxHandler) SetDataAuthority(c *gin.Context) {
	var auth system.Role
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.SetDataAuthority(auth); err != nil {
		zvar.Log.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}
