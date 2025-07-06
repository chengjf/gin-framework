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
		commonGroup.POST("/api/v1/auth/login", backend.Auth.Login)
		commonGroup.POST("/api/v1/auth/logout", backend.Auth.Logout)
		// 上传附件
		commonGroup.POST("/attachment/upload", backend.Attachment.Upload)
	}
	return commonGroup
}
