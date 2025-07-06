package backend

import (
	"gin-framework/app/controller/base"

	"gin-framework/app/service/backend"

	"gin-framework/pkg/response"

	"gin-framework/types/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
}

var User = UserController{}

// Index 获取列表
func (c *UserController) Index(ctx *gin.Context) {
	var requestParams user.IndexRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetIndex(requestParams, ctx)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}

// List 获取列表
func (c *UserController) List(ctx *gin.Context) {
	var requestParams user.IndexRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetList(requestParams, ctx)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}

func (c *UserController) Create(ctx *gin.Context) {
	var requestParams user.UserCreateRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	err := backend.User.Create(requestParams, ctx)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "创建成功", "")
}
