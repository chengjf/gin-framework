package base

import (
	"errors"
	"fmt"
	"gin-framework/pkg/response"
	"net/http"
	"strings"

	"gin-framework/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
func (c *Controller) ValidateReqParams(ctx *gin.Context, requestParams any) error {
	var err error

	if requestParams == nil {
		return errors.New("参数结构体不能为nil")
	}
	// 获取Content-Type并清理可能的附加参数
	contentType := strings.SplitN(ctx.ContentType(), ";", 2)[0]

	switch contentType {
	case gin.MIMEJSON:
		err = ctx.ShouldBindJSON(requestParams)
	case gin.MIMEXML, gin.MIMEXML2:
		err = ctx.ShouldBindXML(requestParams)
	case gin.MIMEPOSTForm:
		err = ctx.ShouldBind(requestParams)
	case gin.MIMEMultipartPOSTForm:
		err = ctx.ShouldBindWith(requestParams, binding.FormMultipart)
	default:
		if ctx.Request.Method == http.MethodGet || ctx.Request.Method == http.MethodDelete {
			err = ctx.ShouldBindQuery(requestParams)
			if uriErr := ctx.ShouldBindUri(requestParams); uriErr != nil && !shouldSkipError(uriErr) {
				if err == nil {
					err = uriErr
				} else {
					err = fmt.Errorf("%w; URI参数错误: %v", err, uriErr)
				}
			}
		} else {
			err = ctx.ShouldBind(requestParams)
		}
	}
	if err != nil {
		translate := validator.Translate(err)
		return errors.New(translate[0])
	}
	return nil
}

// 处理跳过错误的通用方法
func shouldSkipError(err error) bool {
	return err.Error() == "skip" || // 兼容旧版
		err.Error() == "skip binding" // 某些版本的错误信息
}
