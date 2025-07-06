package router

import (
	"fmt"

	"gin-framework/config"
	"gin-framework/global"
	"gin-framework/middleware"
	"gin-framework/pkg/response"
	"gin-framework/router/routes"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	gin.SetMode(global.Cfg.Server.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	// [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	r.SetTrustedProxies(config.Whitelist)
	// cors
	r.Use(middleware.CorsAuth())
	// header add X-Request-Id
	r.Use(requestid.New())
	r.Use(middleware.RequestIdAuth())
	// 404 not found
	r.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		response.NotFoundException(ctx, fmt.Sprintf("%s %s not found", method, path))
	})

	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			middleware.IpAuth(),
		}
		commonGroup = r.Group("/", publicMiddleware...)
		authGroup   = r.Group("/", append(publicMiddleware, middleware.LoginAuth())...)
	)
	// 公用组
	routes.InitCommonGroup(commonGroup)
	routes.InitUserGroup(commonGroup)
	// 后台组
	routes.InitBackendGroup(authGroup)
	// 赋给全局
	global.Router = r
	return r
}
