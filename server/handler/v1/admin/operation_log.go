package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/handler/middleware"
	"project/model/system"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type operationRecordHandler struct {
	service *service.OperationRecordService
}

func OperationRecordHandler() *operationRecordHandler {
	return &operationRecordHandler{
		service: &service.OperationRecordService{},
	}
}

func (h *operationRecordHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createSysOperationRecord", h.CreateSysOperationRecord)             // 新建SysOperationRecord
		apiRouter.DELETE("deleteSysOperationRecord", h.DeleteSysOperationRecord)           // 删除SysOperationRecord
		apiRouter.DELETE("deleteSysOperationRecordByIds", h.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		apiRouter.GET("findSysOperationRecord", h.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		apiRouter.GET("getSysOperationRecordList", h.GetSysOperationRecordList)            // 获取SysOperationRecord列表
	}
}

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (h *operationRecordHandler) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := h.service.CreateSysOperationRecord(sysOperationRecord); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (h *operationRecordHandler) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := h.service.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (h *operationRecordHandler) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := h.service.DeleteSysOperationRecordByIds(IDS); err != nil {
		zvar.Log.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (h *operationRecordHandler) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindQuery(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysOperationRecord := h.service.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysOperationRecordSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (h *operationRecordHandler) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := h.service.GetSysOperationRecordInfoList(pageInfo); err != nil {
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
