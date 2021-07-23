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

type dictionaryHandler struct {
	service *service.DictionaryService
}

func NewDictionaryHandler() *dictionaryHandler {
	return &dictionaryHandler{
		service: &service.DictionaryService{},
	}
}

func (h *dictionaryHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("sysDictionary").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createSysDictionary", h.CreateSysDictionary)   // 新建SysDictionary
		apiRouter.DELETE("deleteSysDictionary", h.DeleteSysDictionary) // 删除SysDictionary
		apiRouter.PUT("updateSysDictionary", h.UpdateSysDictionary)    // 更新SysDictionary
		apiRouter.GET("findSysDictionary", h.FindSysDictionary)        // 根据ID获取SysDictionary
		apiRouter.GET("getSysDictionaryList", h.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionary/createSysDictionary [post]
func (h *dictionaryHandler) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.CreateSysDictionary(dictionary); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteSysDictionary [delete]
func (h *dictionaryHandler) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.DeleteSysDictionary(dictionary); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateSysDictionary [put]
func (h *dictionaryHandler) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := h.service.UpdateSysDictionary(&dictionary); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findSysDictionary [get]
func (h *dictionaryHandler) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindQuery(&dictionary)
	if err, sysDictionary := h.service.GetSysDictionary(dictionary.Type, dictionary.ID); err != nil {
		zvar.Log.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
	}
}

// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionarySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getSysDictionaryList [get]
func (h *dictionaryHandler) GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := h.service.GetSysDictionaryInfoList(pageInfo); err != nil {
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
