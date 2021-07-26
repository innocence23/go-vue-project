package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/entity"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type roleHandler struct {
	roleService *service.RoleService
	rbacService *service.RbacService
}

func NewRoleHandler() *roleHandler {
	return &roleHandler{
		roleService: &service.RoleService{},
		rbacService: &service.RbacService{},
	}
}

func (h *roleHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("role")
	{
		apiRouter.POST("list", h.list)
		apiRouter.POST("create", h.create)
		apiRouter.PUT("update", h.update)
		apiRouter.POST("delete", h.delete)

		apiRouter.POST("setRoleUser", h.setRoleUser)       // 设置角色资源权限
		apiRouter.POST("setRolePermission", h.setRoleUser) // 设置角色资源权限
		apiRouter.POST("setRoleMenu", h.setRoleUser)       // 设置角色资源权限
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/list"] = zvar.RouteInfo{Group: "role", Name: "角色列表"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/create"] = zvar.RouteInfo{Group: "role", Name: "创建角色"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/create"] = zvar.RouteInfo{Group: "role", Name: "更新角色"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/delete"] = zvar.RouteInfo{Group: "role", Name: "删除角色"}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/setRoleUser"] = zvar.RouteInfo{Group: "role", Name: "用户添加角色"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/setRolePermission"] = zvar.RouteInfo{Group: "role", Name: "显示菜单"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/role/setRoleMenu"] = zvar.RouteInfo{Group: "role", Name: "菜单树"}
}

// @Tags Role
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /role/list [post]
func (h *roleHandler) list(c *gin.Context) {
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

// @Tags Role
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Role true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /role/create [post]
func (h *roleHandler) create(c *gin.Context) {
	var req entity.Role
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.roleService.Create(req); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Role
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Role true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /role/update [post]
func (h *roleHandler) update(c *gin.Context) {
	var req entity.Role
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.roleService.Update(req); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Role
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Role true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /role/delete [post]
func (h *roleHandler) delete(c *gin.Context) {
	var req request.IdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.roleService.Delete(req.ID); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Role
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Role true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /role/setRoleUser [post]
func (h *roleHandler) setRoleUser(c *gin.Context) {
	var req request.SetUserRole
	_ = c.ShouldBindJSON(&req)
	if UserVerifyErr := utils.Verify(req, utils.SetUserRoleorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if _, err := h.rbacService.AddRoleForUser(cast.ToString(req.UserID), cast.ToString(req.RoleID)); err != nil {
		zvar.Log.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// // @Tags Authority
// // @Summary 设置角色资源权限
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data body system.Role true "设置角色资源权限"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// // @Router /authority/setDataAuthority [post]
// func (h *rolexxxHandler) SetDataAuthority(c *gin.Context) {
// 	var auth system.Role
// 	_ = c.ShouldBindJSON(&auth)
// 	if err := utils.Verify(auth, utils.RoleIdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := h.service.SetDataAuthority(auth); err != nil {
// 		zvar.Log.Error("设置失败!", zap.Any("err", err))
// 		response.FailWithMessage("设置失败"+err.Error(), c)
// 	} else {
// 		response.OkWithMessage("设置成功", c)
// 	}
// }
