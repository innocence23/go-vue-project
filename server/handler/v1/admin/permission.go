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

type permissionHandler struct {
	service *service.ApiService
}

func NewPermissionHandler() *permissionHandler {
	return &permissionHandler{
		service: &service.ApiService{},
	}
}

func (h *permissionHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("api").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createApi", h.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", h.DeleteApi)               // 删除Api
		apiRouter.POST("getApiList", h.GetApiList)             // 获取Api列表
		apiRouter.POST("getApiById", h.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", h.UpdateApi)               // 更新api
		apiRouter.POST("getAllApis", h.GetAllApis)             // 获取所有api
		apiRouter.DELETE("deleteApisByIds", h.DeleteApisByIds) // 删除选中api
	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func (h *permissionHandler) CreateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.CreateApi(api); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func (h *permissionHandler) DeleteApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api.Method, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.DeleteApi(api); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (h *permissionHandler) GetApiList(c *gin.Context) {
	var pageInfo request.SearchApiParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := h.service.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (h *permissionHandler) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, api := h.service.GetApiById(idInfo.ID)
	if err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.SysAPIResponse{Api: api}, c)
	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func (h *permissionHandler) UpdateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.UpdateApi(api); err != nil {
		zvar.Log.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func (h *permissionHandler) GetAllApis(c *gin.Context) {
	if err, apis := h.service.GetAllApis(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysAPIListResponse{Apis: apis}, "获取成功", c)
	}
}

// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
func (h *permissionHandler) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := h.service.DeleteApisByIds(ids); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
