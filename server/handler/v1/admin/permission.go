package admin

import (
	"project/global"
	"project/handler/middleware"
	"project/model/common/request"
	"project/model/common/response"
	"project/model/system"
	systemReq "project/model/system/request"
	systemRes "project/model/system/response"
	"project/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type permissionHandler struct {
}

func NewPermissionHandler() *permissionHandler {
	return &permissionHandler{}
}

func (s *permissionHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("api").Use(middleware.OperationRecord())
	{
		apiRouter.POST("createApi", s.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", s.DeleteApi)               // 删除Api
		apiRouter.POST("getApiList", s.GetApiList)             // 获取Api列表
		apiRouter.POST("getApiById", s.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", s.UpdateApi)               // 更新api
		apiRouter.POST("getAllApis", s.GetAllApis)             // 获取所有api
		apiRouter.DELETE("deleteApisByIds", s.DeleteApisByIds) // 删除选中api
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
func (s *permissionHandler) CreateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
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
func (s *permissionHandler) DeleteApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.DeleteApi(api); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
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
// @Param data body systemReq.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (s *permissionHandler) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
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
func (s *permissionHandler) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, api := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(systemRes.SysAPIResponse{Api: api}, c)
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
func (s *permissionHandler) UpdateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.UpdateApi(api); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
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
func (s *permissionHandler) GetAllApis(c *gin.Context) {
	if err, apis := apiService.GetAllApis(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "获取成功", c)
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
func (s *permissionHandler) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := apiService.DeleteApisByIds(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
