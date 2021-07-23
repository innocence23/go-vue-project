package autocode

import (
	"project/handler/middleware"
	"project/model/autocode"
	autocodeReq "project/model/autocode/request"
	"project/model/common/response"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var autoCodeExampleService = service.ServiceGroupApp.AutoCodeServiceGroup.AutoCodeExampleService

type autoCodeExampleHandler struct {
}

func NewAutoCodeExampleHandler() *autoCodeExampleHandler {
	return &autoCodeExampleHandler{}
}

func (h *autoCodeExampleHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("autoCodeExample").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createSysAutoCodeExample", h.CreateAutoCodeExample)   // 新建AutoCodeExample
		apiRouter.DELETE("deleteSysAutoCodeExample", h.DeleteAutoCodeExample) // 删除AutoCodeExample
		apiRouter.PUT("updateSysAutoCodeExample", h.UpdateAutoCodeExample)    // 更新AutoCodeExample
		apiRouter.GET("findSysAutoCodeExample", h.FindAutoCodeExample)        // 根据ID获取AutoCodeExample
		apiRouter.GET("getSysAutoCodeExampleList", h.GetAutoCodeExampleList)  // 获取AutoCodeExample列表
	}
}

// @Tags AutoCodeExample
// @Summary 创建AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "AutoCodeExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCodeExample/createAutoCodeExample [post]
func (h *autoCodeExampleHandler) CreateAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.CreateAutoCodeExample(autoCodeExample); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 删除AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "AutoCodeExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /autoCodeExample/deleteAutoCodeExample [delete]
func (h *autoCodeExampleHandler) DeleteAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.DeleteAutoCodeExample(autoCodeExample); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 更新AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "更新AutoCodeExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /autoCodeExample/updateAutoCodeExample [put]
func (h *autoCodeExampleHandler) UpdateAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.UpdateAutoCodeExample(&autoCodeExample); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 用id查询AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "用id查询AutoCodeExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /autoCodeExample/findAutoCodeExample [get]
func (h *autoCodeExampleHandler) FindAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindQuery(&autoCodeExample)
	if err := utils.Verify(autoCodeExample, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, reAutoCodeExample := autoCodeExampleService.GetAutoCodeExample(autoCodeExample.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reAutoCodeExample": reAutoCodeExample}, "查询成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 分页获取AutoCodeExample列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocodeReq.AutoCodeExampleSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCodeExample/getAutoCodeExampleList [get]
func (h *autoCodeExampleHandler) GetAutoCodeExampleList(c *gin.Context) {
	var pageInfo autocodeReq.AutoCodeExampleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := autoCodeExampleService.GetAutoCodeExampleInfoList(pageInfo); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
