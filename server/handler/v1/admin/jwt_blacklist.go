package admin

import (
	"project/dto/response"
	"project/handler/middleware"
	"project/model/system"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type jwtHandler struct {
	service *service.JwtService
}

func NewJwtHandler() *jwtHandler {
	return &jwtHandler{
		service: &service.JwtService{},
	}
}

func (h *jwtHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("jwt").Use(middleware.OperationRecord())
	apiRouter.POST("jsonInBlacklist", h.JsonInBlacklist) // jwt加入黑名单

}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (h *jwtHandler) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := h.service.JsonInBlacklist(jwt); err != nil {
		zvar.Log.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
