package middleware

import (
	"project/model/system/request"
	"project/service"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		_ = success
		//if global.GVA_CONFIG.System.Env == "develop" || success {
		c.Next()
		// } else {
		// 	response.FailWithDetailed(gin.H{}, "权限不足", c)
		// 	c.Abort()
		// 	return
		// }
	}
}
