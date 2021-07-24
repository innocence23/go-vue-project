package middleware

import (
	"project/dto/request"
	"project/service"

	"github.com/gin-gonic/gin"
)

var casbinService = &service.CasbinService{}

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := waitUse.RoleId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		_ = success
		//if zvar.Config.System.Env == "develop" || success {
		c.Next()
		// } else {
		// 	response.FailWithDetailed(gin.H{}, "权限不足", c)
		// 	c.Abort()
		// 	return
		// }
	}
}
