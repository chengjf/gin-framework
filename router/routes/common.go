package routes

import (
	"gin-framework/app/controller/backend"

	"gin-framework/app/controller/common"

	"github.com/gin-gonic/gin"
)

func InitCommonGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	commonGroup := r.Group("")
	{
		// ping
		commonGroup.GET("/ping", common.Common.Ping)
		// 登录
		commonGroup.POST("/user/login", backend.Auth.Login)
		// 上传附件
		commonGroup.POST("/attachment/upload", backend.Attachment.Upload)
	}
	return commonGroup
}
