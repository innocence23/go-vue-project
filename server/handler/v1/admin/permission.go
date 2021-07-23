package admin

import (
	"project/dto/request"
	"project/dto/response"
	"project/model/system"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type permissionHandler struct {
	permService *service.PermissionService
}

func NewPermissionHandler() *permissionHandler {
	return &permissionHandler{
		permService: &service.PermissionService{},
	}
}

func (h *permissionHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("permission")
	{
		apiRouter.POST("create", h.create)              // 创建Api
		apiRouter.POST("delete", h.delete)              // 删除Api
		apiRouter.POST("list", h.list)                  // 获取Api列表
		apiRouter.POST("show", h.show)                  // 获取单条Api消息
		apiRouter.POST("update", h.update)              // 更新api
		apiRouter.POST("listAll", h.listAll)            // 获取所有api
		apiRouter.DELETE("deleteByIds", h.deletesByIds) // 删除选中api
	}
}

// @Tags Permission
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Permission true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func (h *permissionHandler) create(c *gin.Context) {
	var api system.Permission
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.permService.Create(api); err != nil {
		zvar.Log.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Permission
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Permission true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func (h *permissionHandler) delete(c *gin.Context) {
	var req system.Permission
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req.Method, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.permService.Delete(req); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Permission
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (h *permissionHandler) list(c *gin.Context) {
	var pageInfo request.SearchPermissionParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := h.permService.List(pageInfo.Permission, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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

// @Tags Permission
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (h *permissionHandler) show(c *gin.Context) {
	var req request.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, permission := h.permService.Show(req.ID)
	if err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.PermissionResponse{Permission: permission}, c)
	}
}

// @Tags Permission
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Permission true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func (h *permissionHandler) update(c *gin.Context) {
	var api system.Permission
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.permService.Update(api); err != nil {
		zvar.Log.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags Permission
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func (h *permissionHandler) listAll(c *gin.Context) {
	if err, Permissions := h.permService.ListNoLimit(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PermissionListResponse{Permissions: Permissions}, "获取成功", c)
	}
}

// @Tags Permission
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
func (h *permissionHandler) deletesByIds(c *gin.Context) {
	var req request.IdsReq
	_ = c.ShouldBindJSON(&req)
	if err := h.permService.DeleteByIds(req.Ids); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
