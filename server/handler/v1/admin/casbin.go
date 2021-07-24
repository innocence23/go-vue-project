package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type casbinHandler struct {
	casbinService *service.CasbinService
}

func NewCasbinHandler() *casbinHandler {
	return &casbinHandler{
		casbinService: &service.CasbinService{},
	}
}

func (h *casbinHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("casbin")
	apiRouter.POST("update", h.update)
	apiRouter.POST("getPermByRoleId", h.getPermByRoleId)

	zvar.RouteMap = map[string]zvar.RouteInfo{
		"/" + zvar.UrlPrefix + "/casbin/update":          {Group: "casbin", Name: "更新角色权限"},
		"/" + zvar.UrlPrefix + "/casbin/getPermByRoleId": {Group: "casbin", Name: "获取角色权限列表"},
	}
}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/update [post]
func (h *casbinHandler) update(c *gin.Context) {
	var req request.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.casbinService.Update(req.RoleId, req.CasbinInfos); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags req
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPermByRoleId [post]
func (h *casbinHandler) getPermByRoleId(c *gin.Context) {
	var req request.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := h.casbinService.GetPermByRoleId(req.RoleId)
	response.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
