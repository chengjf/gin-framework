package middleware

import (
	"gin-framework/global"
	"gin-framework/pkg/auth"
	"gin-framework/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// LoginAuth 登录中间件
func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(global.Cfg.Jwt.TokenKey)
		if token == "" {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		b := strings.Contains(token, "Bearer")
		if !b {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
		if _, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret); err != nil {
			response.UnauthorizedException(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
