package admin

import (
	"encoding/json"
	"project/dto/request"
	"project/dto/response"
	"project/entity"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type menuHandler struct {
	menuService *service.MenuService
}

func NewMenuHandler() *menuHandler {
	return &menuHandler{
		menuService: &service.MenuService{},
	}
}

func (h *menuHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("menu")
	{
		apiRouter.POST("list", h.list)
		apiRouter.POST("show", h.show)
		apiRouter.POST("create", h.create)
		apiRouter.POST("update", h.update)
		apiRouter.POST("hidden", h.hidden)
		apiRouter.POST("display", h.display)

		apiRouter.POST("treeList", h.treeList)
	}

	//todo
	// zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/getUidMenu"] = zvar.RouteInfo{Group: "menu", Name: "用户菜单"}
	// zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/getRoleMenu"] = zvar.RouteInfo{Group: "menu", Name: "角色菜单"}
	// zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/addRoleMenu"] = zvar.RouteInfo{Group: "menu", Name: "角色添加菜单"}

	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/list"] = zvar.RouteInfo{Group: "menu", Name: "菜单列表"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/show"] = zvar.RouteInfo{Group: "menu", Name: "菜单详情"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/create"] = zvar.RouteInfo{Group: "menu", Name: "新增菜单"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/update"] = zvar.RouteInfo{Group: "menu", Name: "更新菜单"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/hidden"] = zvar.RouteInfo{Group: "menu", Name: "隐藏菜单"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/display"] = zvar.RouteInfo{Group: "menu", Name: "显示菜单"}
	zvar.RouteMap["/"+zvar.UrlPrefix+"/menu/treeList"] = zvar.RouteInfo{Group: "menu", Name: "菜单树"}
}

// @Tags AuthorityMenu    //todo 待改进 []存储 不用中间表）
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/treeList [post]
func (h *menuHandler) treeList(c *gin.Context) {
	str := `{"menus":[{"ID":22,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"https://www.gin-vue-admin.com","name":"https://www.gin-vue-admin.com","hidden":false,"component":"/","sort":0,"meta":{"keepAlive":false,"defaultMenu":false,"title":"官方网站","icon":"s-home","closeTab":false},"authoritys":null,"menuId":"22","children":null,"parameters":[]},{"ID":1,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"dashboard","name":"dashboard","hidden":false,"component":"view/dashboard/index.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"仪表盘","icon":"setting","closeTab":false},"authoritys":null,"menuId":"1","children":null,"parameters":[]},{"ID":3,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"admin","name":"superAdmin","hidden":false,"component":"view/superAdmin/index.vue","sort":3,"meta":{"keepAlive":false,"defaultMenu":false,"title":"超级管理员","icon":"user-solid","closeTab":false},"authoritys":null,"menuId":"3","children":[{"ID":4,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"authority","name":"authority","hidden":false,"component":"view/superAdmin/authority/authority.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"角色管理","icon":"s-custom","closeTab":false},"authoritys":null,"menuId":"4","children":null,"parameters":[]},{"ID":19,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-24T15:12:17+08:00","parentId":"3","path":"dictionaryDetail/:id","name":"dictionaryDetail","hidden":false,"component":"view/superAdmin/dictionary/sysDictionaryDetail.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"字典详情","icon":"s-order","closeTab":false},"authoritys":null,"menuId":"19","children":null,"parameters":[]},{"ID":5,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"menu","name":"menu","hidden":false,"component":"view/superAdmin/menu/menu.vue","sort":2,"meta":{"keepAlive":true,"defaultMenu":false,"title":"菜单管理","icon":"s-order","closeTab":false},"authoritys":null,"menuId":"5","children":null,"parameters":[]},{"ID":6,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"api","name":"api","hidden":false,"component":"view/superAdmin/api/api.vue","sort":3,"meta":{"keepAlive":true,"defaultMenu":false,"title":"api管理","icon":"s-platform","closeTab":false},"authoritys":null,"menuId":"6","children":null,"parameters":[]},{"ID":7,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"user","name":"user","hidden":false,"component":"view/superAdmin/user/user.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"用户管理","icon":"coordinate","closeTab":false},"authoritys":null,"menuId":"7","children":null,"parameters":[]},{"ID":18,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"dictionary","name":"dictionary","hidden":false,"component":"view/superAdmin/dictionary/sysDictionary.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"字典管理","icon":"notebook-2","closeTab":false},"authoritys":null,"menuId":"18","children":null,"parameters":[]},{"ID":20,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"3","path":"operation","name":"operation","hidden":false,"component":"view/superAdmin/operation/sysOperationRecord.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"操作历史","icon":"time","closeTab":false},"authoritys":null,"menuId":"20","children":null,"parameters":[]}],"parameters":[]},{"ID":8,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"person","name":"person","hidden":true,"component":"view/person/person.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"个人信息","icon":"message-solid","closeTab":false},"authoritys":null,"menuId":"8","children":null,"parameters":[]},{"ID":14,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"systemTools","name":"systemTools","hidden":false,"component":"view/systemTools/index.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"系统工具","icon":"s-cooperation","closeTab":false},"authoritys":null,"menuId":"14","children":[{"ID":25,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"14","path":"autoCodeEdit/:id","name":"autoCodeEdit","hidden":true,"component":"view/systemTools/autoCode/index.vue","sort":0,"meta":{"keepAlive":false,"defaultMenu":false,"title":"自动化代码（复用）","icon":"s-finance","closeTab":false},"authoritys":null,"menuId":"25","children":null,"parameters":[]},{"ID":15,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"14","path":"autoCode","name":"autoCode","hidden":false,"component":"view/systemTools/autoCode/index.vue","sort":1,"meta":{"keepAlive":true,"defaultMenu":false,"title":"代码生成器","icon":"cpu","closeTab":false},"authoritys":null,"menuId":"15","children":null,"parameters":[]},{"ID":24,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"14","path":"autoCodeAdmin","name":"autoCodeAdmin","hidden":false,"component":"view/systemTools/autoCodeAdmin/index.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"自动化代码管理","icon":"s-finance","closeTab":false},"authoritys":null,"menuId":"24","children":null,"parameters":[]},{"ID":16,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"14","path":"formCreate","name":"formCreate","hidden":false,"component":"view/systemTools/formCreate/index.vue","sort":2,"meta":{"keepAlive":true,"defaultMenu":false,"title":"表单生成器","icon":"magic-stick","closeTab":false},"authoritys":null,"menuId":"16","children":null,"parameters":[]}],"parameters":[]},{"ID":9,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"example","name":"example","hidden":false,"component":"view/example/index.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"示例文件","icon":"s-management","closeTab":false},"authoritys":null,"menuId":"9","children":[{"ID":10,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"9","path":"excel","name":"excel","hidden":false,"component":"view/example/excel/excel.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"excel导入导出","icon":"s-marketing","closeTab":false},"authoritys":null,"menuId":"10","children":null,"parameters":[]},{"ID":11,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"9","path":"upload","name":"upload","hidden":false,"component":"view/example/upload/upload.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"媒体库（上传下载）","icon":"upload","closeTab":false},"authoritys":null,"menuId":"11","children":null,"parameters":[]},{"ID":12,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"9","path":"breakpoint","name":"breakpoint","hidden":false,"component":"view/example/breakpoint/breakpoint.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"断点续传","icon":"upload","closeTab":false},"authoritys":null,"menuId":"12","children":null,"parameters":[]},{"ID":21,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"9","path":"simpleUploader","name":"simpleUploader","hidden":false,"component":"view/example/simpleUploader/simpleUploader","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"断点续传（插件版）","icon":"upload","closeTab":false},"authoritys":null,"menuId":"21","children":null,"parameters":[]},{"ID":13,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"9","path":"customer","name":"customer","hidden":false,"component":"view/example/customer/customer.vue","sort":7,"meta":{"keepAlive":false,"defaultMenu":false,"title":"客户列表（资源示例）","icon":"s-custom","closeTab":false},"authoritys":null,"menuId":"13","children":null,"parameters":[]}],"parameters":[]},{"ID":23,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"state","name":"state","hidden":false,"component":"view/system/state.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"服务器状态","icon":"cloudy","closeTab":false},"authoritys":null,"menuId":"23","children":null,"parameters":[]},{"ID":2,"CreatedAt":"2021-07-22T18:48:47+08:00","UpdatedAt":"2021-07-22T18:48:47+08:00","parentId":"0","path":"about","name":"about","hidden":false,"component":"view/about/index.vue","sort":7,"meta":{"keepAlive":false,"defaultMenu":false,"title":"关于我们","icon":"info","closeTab":false},"authoritys":null,"menuId":"2","children":null,"parameters":[]}]}`
	var data map[string]interface{}
	json.Unmarshal([]byte(str), &data)

	response.OkWithDetailed(data, "获取成功", c)
	return
	//todo
	if err, menus := h.menuService.GetMenuTree(utils.GetUserRoleId(c)[0]); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []entity.Menu{}
		}
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/list [post]
func (h *menuHandler) list(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menuList, err := h.menuService.TreeList(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdReq true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/show [post]
func (h *menuHandler) show(c *gin.Context) {
	var req request.IdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := h.menuService.Find(req.ID); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Menu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/create [post]
func (h *menuHandler) create(c *gin.Context) {
	var req entity.Menu
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.menuService.Create(req); err != nil {
		zvar.Log.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body entity.Menu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/update [post]
func (h *menuHandler) update(c *gin.Context) {
	var req entity.Menu
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.menuService.Update(req); err != nil {
		zvar.Log.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Menu
// @Summary 隐藏菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdReq true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/hidden [post]
func (h *menuHandler) hidden(c *gin.Context) {
	var req request.IdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.menuService.Hidden(req.ID); err != nil {
		zvar.Log.Error("隐藏失败!", zap.Any("err", err))
		response.FailWithMessage("隐藏失败", c)
	} else {
		response.OkWithMessage("隐藏成功", c)
	}
}

// @Tags Menu
// @Summary 显示菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdReq true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/display [post]
func (h *menuHandler) display(c *gin.Context) {
	var req request.IdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := h.menuService.Display(req.ID); err != nil {
		zvar.Log.Error("显示失败!", zap.Any("err", err))
		response.FailWithMessage("显示失败", c)
	} else {
		response.OkWithMessage("显示成功", c)
	}
}
