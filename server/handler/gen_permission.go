package handler

import (
	"project/entity"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
)

func autoGenPermission(routes gin.RoutesInfo) {
	for _, route := range routes {
		permission := entity.Permission{
			Path:        route.Path,
			Description: zvar.RouteMap[route.Path].Name,
			Group:    zvar.RouteMap[route.Path].Group,
			Method:      route.Method,
		}
		(&service.PermissionService{}).Create(permission)
	}
}
