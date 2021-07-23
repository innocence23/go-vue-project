package admin

import (
	"project/handler/middleware"
	"project/model/common/response"
	"project/zvar"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type emailHandler struct {
}

func NewEmailHandler() *emailHandler {
	return &emailHandler{}
}

func (e *emailHandler) Router(router *gin.RouterGroup) {
	apiRouter := router.Group("email").Use(middleware.OperationRecord())
	apiRouter.POST("emailTest", e.EmailTest) // 发送测试邮件
}

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (e *emailHandler) EmailTest(c *gin.Context) {
	if err := emailService.EmailTest(); err != nil {
		zvar.Log.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
