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

type dictHandler struct {
	service *service.DictService
}

func NewDictHandler() *dictHandler {
	return &dictHandler{
		service: &service.DictService{},
	}
}

func (h *dictHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("dict").Use(middleware.OperationRecord())
	{
		apiRouter.POST("create", h.create)
		apiRouter.DELETE("delete", h.delete)
		apiRouter.PUT("update", h.update)
		apiRouter.GET("show", h.show)
		apiRouter.GET("list", h.list)
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/dict/create"] = zvar.RouteInfo{Group: "dict", Name: "创建字典"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/dict/delete"] = zvar.RouteInfo{Group: "dict", Name: "删除字典"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/dict/update"] = zvar.RouteInfo{Group: "dict", Name: "更新字典"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/dict/show"] = zvar.RouteInfo{Group: "dict", Name: "字典详情"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/dict/list"] = zvar.RouteInfo{Group: "dict", Name: "字典列表"}
}

// @Tags Dict
// @Summary 创建Dict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Dict true "Dict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dict/create [post]
func (h *dictHandler) create(c *gin.Context) {
	var dictionary entity.Dict
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.CreateDict(dictionary); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Dict
// @Summary 删除Dict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Dict true "Dict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dict/delete [delete]
func (h *dictHandler) delete(c *gin.Context) {
	var dictionary entity.Dict
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.DeleteDict(dictionary); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Dict
// @Summary 更新Dict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Dict true "Dict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dict/update [put]
func (h *dictHandler) update(c *gin.Context) {
	var dictionary entity.Dict
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.UpdateDict(&dictionary); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Dict
// @Summary 用id查询Dict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Dict true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dict/show [get]
func (h *dictHandler) show(c *gin.Context) {
	var dictionary entity.Dict
	_ = c.ShouldBindQuery(&dictionary)
	if err, dict := h.service.GetDict(dictionary.Type, dictionary.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"redict": dict}, "查询成功", c)
	}
}

// @Tags Dict
// @Summary 分页获取Dict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict/list [get]
func (h *dictHandler) list(c *gin.Context) {
	var pageInfo request.DictSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := h.service.GetDictInfoList(pageInfo); err != nil {
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
