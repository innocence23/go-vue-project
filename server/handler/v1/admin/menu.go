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

type menuHandler struct {
	service   *service.MenuService
	serviceBm *service.BaseMenuService
}

func NewMenuHandler() *menuHandler {
	return &menuHandler{
		service:   &service.MenuService{},
		serviceBm: &service.BaseMenuService{},
	}
}

func (h *menuHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("menu").Use(middleware.OperationRecord())
	{
		apiRouter.POST("tree", h.GetMenu)
		apiRouter.POST("list", h.GetMenuList)                  // 分页获取基础menu列表
		apiRouter.POST("addBaseMenu", h.AddBaseMenu)           // 新增菜单
		apiRouter.POST("getBaseMenuTree", h.GetBaseMenuTree)   // 获取用户动态路由
		apiRouter.POST("addMenuAuthority", h.AddMenuAuthority) //	增加menu和角色关联关系
		apiRouter.POST("getMenuAuthority", h.GetMenuAuthority) // 获取指定角色menu
		apiRouter.POST("deleteBaseMenu", h.DeleteBaseMenu)     // 删除菜单
		apiRouter.POST("updateBaseMenu", h.UpdateBaseMenu)     // 更新菜单
		apiRouter.POST("getBaseMenuById", h.GetBaseMenuById)   // 根据id获取菜单
	}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/tree"] = zvar.RouteInfo{Group: "menu", Name: "菜单树"}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func (h *menuHandler) GetMenu(c *gin.Context) {
	if err, menus := h.service.GetMenuTree(utils.GetUserRoleId(c)); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func (h *menuHandler) GetBaseMenuTree(c *gin.Context) {
	if err, menus := h.service.GetBaseMenuTree(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func (h *menuHandler) AddMenuAuthority(c *gin.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.AddMenuAuthority(authorityMenu.Menus, authorityMenu.RoleId); err != nil {
		zvar.Log.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetRoleId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (h *menuHandler) GetMenuAuthority(c *gin.Context) {
	var param request.GetRoleId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := h.service.GetMenuAuthority(&param); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (h *menuHandler) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.service.AddBaseMenu(menu); err != nil {
		zvar.Log.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (h *menuHandler) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.serviceBm.DeleteBaseMenu(menu.ID); err != nil {
		zvar.Log.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (h *menuHandler) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.serviceBm.UpdateBaseMenu(menu); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (h *menuHandler) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := h.serviceBm.GetBaseMenuById(idInfo.ID); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (h *menuHandler) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menuList, total := h.service.GetInfoList(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
