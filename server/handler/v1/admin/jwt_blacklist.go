package admin

import (
	"project/dto/response"
	"project/model/system"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type jwtHandler struct {
	jwtService *service.JwtService
}

func NewJwtHandler() *jwtHandler {
	return &jwtHandler{
		jwtService: &service.JwtService{},
	}
}

func (h *jwtHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("jwt")
	apiRouter.POST("inBlacklist", h.inBlacklist) // jwt加入黑名单

}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/inBlacklist [post]
func (h *jwtHandler) inBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := h.jwtService.InBlacklist(jwt); err != nil {
		zvar.Log.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
