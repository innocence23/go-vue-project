package utils

import (
	"fmt"
	"project/dto/request"
	"project/zvar"

	"github.com/gin-gonic/gin"
)

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		zvar.Log.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		zvar.Log.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID.String()
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserRoleId(c *gin.Context) string {
	claims, _ := c.Get("claims")
	fmt.Printf("=================%+v", claims)
	if claims, exists := c.Get("claims"); !exists {
		zvar.Log.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.RoleId
	}
}
