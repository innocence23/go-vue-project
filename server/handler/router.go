package handler

import (
	"net/http"
	_ "project/docs"
	"project/handler/middleware"
	"project/handler/v1/admin"
	"project/zvar"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(zvar.Config.Local.Path, http.Dir(zvar.Config.Local.Path)) // 为用户头像和文件提供静态地址
	zvar.Log.Info("use middleware logger")

	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	zvar.Log.Info("use middleware cors")
	Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	zvar.Log.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	Router.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	//获取路由组实例
	gRouter := Router.Group(zvar.UrlPrefix)
	admin.NewBaseHandler().Router(gRouter) // 登陆及验证码

	gRouter.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//.Use(middleware.OperationRecord())
	{
		admin.NewUserHandler().Router(gRouter)       // 用户路由
		admin.NewPermissionHandler().Router(gRouter) // 注册功能api路由
		admin.NewJwtHandler().Router(gRouter)        // jwt相关路由
		admin.NewRoleHandler().Router(gRouter)       // 注册角色路由
		admin.NewMenuHandler().Router(gRouter)       // 注册menu路由
		admin.NewSysHandler().Router(gRouter)        // system相关路由
		admin.NewEmailHandler().Router(gRouter)      // 邮件相关路由
		admin.NewDictHandler().Router(gRouter)       // 字典管理
		admin.NewDictDetailHandler().Router(gRouter) // 字典详情管理

		//admin.NewCasbinHandler().Router(gRouter)     // 权限相关路由
		//admin.NewAutoHandler().Router(gRouter)             // 创建自动化代码
		admin.OperationRecordHandler().Router(gRouter) // 操作记录

	}

	//自动生成路由权限
	appRouters := Router.Routes()
	autoGenPermission(appRouters)

	zvar.Log.Info("router register success")
	return Router
}
