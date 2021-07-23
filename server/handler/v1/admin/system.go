package admin

import (
	"project/dto/response"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type machineHandler struct {
	machineService *service.MachineService
}

func NewSysHandler() *machineHandler {
	return &machineHandler{
		machineService: &service.MachineService{},
	}
}

func (h *machineHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("machine")
	apiRouter.GET("info", h.info) // 获取服务器信息
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/info [post]
func (h *machineHandler) info(c *gin.Context) {
	if server, err := h.machineService.GetServerInfo(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
