package base

import (
	"errors"
	"gin-framework/pkg/response"
	"net/http"

	"gin-framework/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

var Base = Controller{}

func (c *Controller) Index(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "index", "")
}

func (c *Controller) Create(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "create", "")
}

func (c *Controller) Delete(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "delete", "")
}

func (c *Controller) Update(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "update", "")
}

func (c *Controller) View(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "view", "")
}

// ValidateReqParams 验证请求参数
func (c *Controller) ValidateReqParams(ctx *gin.Context, requestParams interface{}) error {
	var err error
	switch ctx.ContentType() {
	case "application/json":
		err = ctx.ShouldBindJSON(requestParams)
	case "application/xml":
		err = ctx.ShouldBindXML(requestParams)
	case "":
		err = ctx.ShouldBindUri(requestParams)
		err = ctx.ShouldBindQuery(requestParams)
	default:
		err = ctx.ShouldBind(requestParams)
	}
	if err != nil {
		translate := validator.Translate(err)
		return errors.New(translate[0])
	}
	return nil
}
