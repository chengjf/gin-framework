package frontend

import (
	"net/http"

	"gin-framework/app/controller/base"

	"gin-framework/app/controller/frontend"

	"gin-framework/app/service/frontend"

	"gin-framework/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*base.Controller
}

var User = &UserController{}

// IndexRequest 获取用户列表请求参数
type CreateRequest struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func (c *UserController) Create(ctx *gin.Context) {
	var createRequest CreateRequest
	if err := c.ValidateReqParams(ctx, &createRequest); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	frontend.User.Create(createRequest)
	response.ResponseJson(ctx, http.StatusOK, response.Success, "create", "")
}
