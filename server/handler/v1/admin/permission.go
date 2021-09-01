package admin

import (
	"project/dto/request"
	"project/dto/response"
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
		apiRouter.POST("list", h.list)
		apiRouter.POST("show", h.show)
		apiRouter.POST("listAll", h.listAll)
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/permission/show"] = zvar.RouteInfo{Group: "permission", Name: "权限详情"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/permission/listAll"] = zvar.RouteInfo{Group: "permission", Name: "权限列表ALL"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/permission/list"] = zvar.RouteInfo{Group: "permission", Name: "权限列表"}
}

// @Tags Permission
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchPermissionParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (h *permissionHandler) list(c *gin.Context) {
	var pageInfo request.SearchPermissionParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := h.permService.List(pageInfo.Permission, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
// @Param data body request.IdReq true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (h *permissionHandler) show(c *gin.Context) {
	var req request.IdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	permission, err := h.permService.Show(req.ID)
	if err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.PermissionResponse{Permission: permission}, c)
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
	if Permissions, err := h.permService.ListNoLimit(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PermissionListResponse{Permissions: Permissions}, "获取成功", c)
	}
}
