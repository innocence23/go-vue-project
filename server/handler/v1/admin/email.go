package admin

import (
	"project/dto/response"
	"project/service"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type emailHandler struct {
	emailService *service.EmailService
}

func NewEmailHandler() *emailHandler {
	return &emailHandler{
		emailService: &service.EmailService{},
	}
}

func (h *emailHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("email")
	apiRouter.POST("test", h.emailTest)

	zvar.RouteMap["/"+zvar.UrlPrefix+"/email/test"] = zvar.RouteInfo{Group: "email", Name: "发送测试邮件"}
}

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/test [post]
func (h *emailHandler) emailTest(c *gin.Context) {
	if err := h.emailService.EmailTest(); err != nil {
		zvar.Log.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
