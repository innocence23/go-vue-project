package admin

import (
	"project/handler/middleware"
	"project/model/common/response"
	"project/model/system"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type jwtHandler struct {
}

func NewJwtHandler() *jwtHandler {
	return &jwtHandler{}
}

func (j *jwtHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("jwt").Use(middleware.OperationRecord())
	apiRouter.POST("jsonInBlacklist", j.JsonInBlacklist) // jwt加入黑名单

}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (j *jwtHandler) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		zvar.Log.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
