package admin

import (
	"project/dto/response"
	"project/handler/middleware"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type sysHandler struct {
	service *service.MachineService
}

func NewSysHandler() *sysHandler {
	return &sysHandler{
		service: &service.MachineService{},
	}
}

func (h *sysHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("system").Use(middleware.OperationRecord())
	apiRouter.POST("getServerInfo", h.GetServerInfo) // 获取服务器信息
	apiRouter.POST("reloadSystem", h.ReloadSystem)   // 重启服务
}

// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"重启系统成功"}"
// @Router /system/reloadSystem [post]
func (h *sysHandler) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		zvar.Log.Error("重启系统失败!", zap.Any("err", err))
		response.FailWithMessage("重启系统失败", c)
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getServerInfo [post]
func (h *sysHandler) GetServerInfo(c *gin.Context) {
	if server, err := h.service.GetServerInfo(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
