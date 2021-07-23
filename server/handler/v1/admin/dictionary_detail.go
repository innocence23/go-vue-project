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

type dictionaryDetailHandler struct {
	service *service.DictionaryDetailService
}

func NewDictionaryDetailHandler() *dictionaryDetailHandler {
	return &dictionaryDetailHandler{
		service: &service.DictionaryDetailService{},
	}
}

func (h *dictionaryDetailHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createSysDictionaryDetail", h.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		apiRouter.DELETE("deleteSysDictionaryDetail", h.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		apiRouter.PUT("updateSysDictionaryDetail", h.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
		apiRouter.GET("findSysDictionaryDetail", h.FindSysDictionaryDetail)        // 根据ID获取SysDictionaryDetail
		apiRouter.GET("getSysDictionaryDetailList", h.GetSysDictionaryDetailList)  // 获取SysDictionaryDetail列表
	}
}

// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (h *dictionaryDetailHandler) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.CreateSysDictionaryDetail(detail); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 删除SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (h *dictionaryDetailHandler) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.DeleteSysDictionaryDetail(detail); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (h *dictionaryDetailHandler) UpdateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.UpdateSysDictionaryDetail(&detail); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (h *dictionaryDetailHandler) FindSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysDictionaryDetail := h.service.GetSysDictionaryDetail(detail.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionaryDetail": resysDictionaryDetail}, "查询成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionaryDetailSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (h *dictionaryDetailHandler) GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := h.service.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
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
