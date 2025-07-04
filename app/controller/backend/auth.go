package backend

import (
	"gin-framework/app/controller/base"

	"gin-framework/app/service/common"

	"gin-framework/pkg/response"

	"gin-framework/types/user"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	base.Controller
}

var Auth = AuthController{}

// Login 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
	var requestParams user.LoginRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	token, err := common.User.Login(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", token)
}
