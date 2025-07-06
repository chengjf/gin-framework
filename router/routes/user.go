package routes

import (
	"gin-framework/app/controller/backend"

	"github.com/gin-gonic/gin"
)

// InitBackendGroup 初始化后台接口路由
func InitUserGroup(r *gin.RouterGroup) gin.IRoutes {
	userGroup := r.Group("api/v1")
	{
		userGroup.GET("/user-info", backend.User.View)
		userGroup.GET("/authority", backend.User.View)

	}
	return userGroup
}
