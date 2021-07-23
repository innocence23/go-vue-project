package admin

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"project/dto/request"
	"project/dto/response"
	"project/model/system"
	"project/service"
	"project/utils"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type autoHandler struct {
	serviceHistory *service.AutoCodeHistoryService
	service        *service.AutoCodeService
}

func NewAutoHandler() *autoHandler {
	return &autoHandler{
		serviceHistory: &service.AutoCodeHistoryService{},
		service:        &service.AutoCodeService{},
	}
}

func (h *autoHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("autoCode")
	{
		apiRouter.POST("delSysHistory", h.DelSysHistory) // 删除回滚记录
		apiRouter.POST("getMeta", h.GetMeta)             // 根据id获取meta信息
		apiRouter.POST("getSysHistory", h.GetSysHistory) // 获取回滚记录分页
		apiRouter.POST("rollback", h.RollBack)           // 回滚
		apiRouter.POST("preview", h.PreviewTemp)         // 获取自动创建代码预览
		apiRouter.POST("createTemp", h.CreateTemp)       // 创建自动化代码
		apiRouter.GET("getTables", h.GetTables)          // 获取对应数据库的表
		apiRouter.GET("getDB", h.GetDB)                  // 获取数据库
		apiRouter.GET("getColumn", h.GetColumn)          // 获取指定表所有字段信息
	}
}

// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "删除回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /autoCode/delSysHistory [post]
func (h *autoHandler) DelSysHistory(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	err := h.serviceHistory.DeletePage(id.ID)
	if err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithMessage("删除成功", c)

}

// @Tags AutoCode
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysAutoHistory true "查询回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getSysHistory [post]
func (h *autoHandler) GetSysHistory(c *gin.Context) {
	var search request.SysAutoHistory
	_ = c.ShouldBindJSON(&search)
	err, list, total := h.serviceHistory.GetSysHistoryPage(search.PageInfo)
	if err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     search.Page,
			PageSize: search.PageSize,
		}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "回滚自动生成代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"回滚成功"}"
// @Router /autoCode/rollback [post]
func (h *autoHandler) RollBack(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if err := h.serviceHistory.RollBack(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回滚成功", c)
}

// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "获取meta信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getMeta [post]
func (h *autoHandler) GetMeta(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if v, err := h.serviceHistory.GetMeta(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(gin.H{"meta": v}, "获取成功", c)
	}

}

// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/preview [post]
func (h *autoHandler) PreviewTemp(c *gin.Context) {
	var ac system.AutoCodeStruct
	_ = c.ShouldBindJSON(&ac)
	if err := utils.Verify(ac, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	autoCode, err := h.service.PreviewTemp(ac)
	if err != nil {
		zvar.Log.Error("预览失败!", zap.Any("err", err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (h *autoHandler) CreateTemp(c *gin.Context) {
	var ac system.AutoCodeStruct
	_ = c.ShouldBindJSON(&ac)
	if err := utils.Verify(ac, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var apiIds []uint
	if ac.AutoCreateApiToSql {
		if ids, err := h.service.AutoCreateApi(&ac); err != nil {
			zvar.Log.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	err := h.service.CreateTemp(ac, apiIds...)
	if err != nil {
		if errors.Is(err, system.AutoMoveErr) {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msgtype", "success")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getTables [get]
func (h *autoHandler) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", zvar.Config.Mysql.Dbname)
	err, tables := h.service.GetTables(dbName)
	if err != nil {
		zvar.Log.Error("查询table失败!", zap.Any("err", err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getDatabase [get]
func (h *autoHandler) GetDB(c *gin.Context) {
	if err, dbs := h.service.GetDB(); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getColumn [get]
func (h *autoHandler) GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", zvar.Config.Mysql.Dbname)
	tableName := c.Query("tableName")
	if err, columns := h.service.GetColumn(tableName, dbName); err != nil {
		zvar.Log.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}
