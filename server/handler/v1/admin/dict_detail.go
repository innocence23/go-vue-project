package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/entity"
	"project/handler/middleware"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type dictDetailHandler struct {
	service *service.DictDetailService
}

func NewDictDetailHandler() *dictDetailHandler {
	return &dictDetailHandler{
		service: &service.DictDetailService{},
	}
}

func (h *dictDetailHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("dict-detail").Use(middleware.OperationRecord())
	{
		apiRouter.POST("create", h.create)
		apiRouter.DELETE("delete", h.delete)
		apiRouter.PUT("update", h.update)
		apiRouter.GET("show", h.show)
		apiRouter.GET("list", h.list)

		zvar.RouteMap["/"+zvar.UrlPrefix+"/dict-detail/create"] = zvar.RouteInfo{Group: "dict-detail", Name: "字典详情创建"}
		zvar.RouteMap["/"+zvar.UrlPrefix+"/dict-detail/delete"] = zvar.RouteInfo{Group: "dict-detail", Name: "字典详情删除"}
		zvar.RouteMap["/"+zvar.UrlPrefix+"/dict-detail/update"] = zvar.RouteInfo{Group: "dict-detail", Name: "字典详情更新"}
		zvar.RouteMap["/"+zvar.UrlPrefix+"/dict-detail/show"] = zvar.RouteInfo{Group: "dict-detail", Name: "字典详情展示"}
		zvar.RouteMap["/"+zvar.UrlPrefix+"/dict-detail/list"] = zvar.RouteInfo{Group: "dict-detail", Name: "字典详情列表"}

	}
}

// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.DictDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dict-detail/createSysDictionaryDetail [post]
func (h *dictDetailHandler) create(c *gin.Context) {
	var detail entity.DictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.CreateDictDetail(detail); err != nil {
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
// @Param data body entity.DictDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dict-detail/deleteSysDictionaryDetail [delete]
func (h *dictDetailHandler) delete(c *gin.Context) {
	var detail entity.DictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.DeleteDictDetail(detail); err != nil {
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
// @Param data body entity.DictDetail true "更新"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dict-detail/updateSysDictionaryDetail [put]
func (h *dictDetailHandler) update(c *gin.Context) {
	var detail entity.DictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := h.service.UpdateDictDetail(&detail); err != nil {
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
// @Param data body entity.DictDetail true "用id查询"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dict-detail/findSysDictionaryDetail [get]
func (h *dictDetailHandler) show(c *gin.Context) {
	var detail entity.DictDetail
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, dictDetail := h.service.GetDictDetail(detail.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dictDetail": dictDetail}, "查询成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictDetailSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict-detail/getSysDictionaryDetailList [get]
func (h *dictDetailHandler) list(c *gin.Context) {
	var pageInfo request.DictDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := h.service.GetDictDetailList(pageInfo); err != nil {
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
