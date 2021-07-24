package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/entity"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type operationRecordHandler struct {
	opLogService *service.OperationLogService
}

func OperationRecordHandler() *operationRecordHandler {
	return &operationRecordHandler{
		opLogService: &service.OperationLogService{},
	}
}

func (h *operationRecordHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("opLog")
	{
		apiRouter.DELETE("delete", h.delete)
		apiRouter.DELETE("deleteByIds", h.deleteByIds)
		apiRouter.GET("list", h.list)
	}
	zvar.RouteMap = map[string]zvar.RouteInfo{
		"/" + zvar.UrlPrefix + "/opLog/delete":      {Group: "opLog", Name: "删除日志"},
		"/" + zvar.UrlPrefix + "/opLog/deleteByIds": {Group: "opLog", Name: "批量删除日志"},
		"/" + zvar.UrlPrefix + "/opLog/list":        {Group: "opLog", Name: "日志列表"},
	}
}

// @Tags OperationLog
// @Summary 删除OperationLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.OperationLog true "OperationLog模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /req/delete [delete]
func (h *operationRecordHandler) delete(c *gin.Context) {
	var req entity.OperationLog
	_ = c.ShouldBindJSON(&req)
	if err := h.opLogService.Delete(req); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags OperationLog
// @Summary 批量删除OperationLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OperationLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /req/deleteByIds [delete]
func (h *operationRecordHandler) deleteByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := h.opLogService.DeleteByIds(IDS); err != nil {
		zvar.Log.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags OperationLogSearch
// @Summary 分页获取OperationLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperationLogSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /req/list [get]
func (h *operationRecordHandler) list(c *gin.Context) {
	var pageInfo request.OperationLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := h.opLogService.List(pageInfo); err != nil {
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
