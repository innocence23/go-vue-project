package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/handler/middleware"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type casbinHandler struct {
	service *service.CasbinService
}

func NewCasbinHandler() *casbinHandler {
	return &casbinHandler{
		service: &service.CasbinService{},
	}
}

func (h *casbinHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("casbin").Use(middleware.OperationRecord())
	apiRouter.POST("updateCasbin", h.UpdateCasbin)
	apiRouter.POST("getPolicyPathByAuthorityId", h.GetPolicyPathByAuthorityId)
}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]
func (h *casbinHandler) UpdateCasbin(c *gin.Context) {
	var req request.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.UpdateCasbin(req.AuthorityId, req.CasbinInfos); err != nil {
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
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (h *casbinHandler) GetPolicyPathByAuthorityId(c *gin.Context) {
	var req request.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := h.service.GetPolicyPathByAuthorityId(req.AuthorityId)
	response.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
