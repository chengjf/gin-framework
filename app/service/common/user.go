package common

import (
	"errors"

	"gin-framework/global"
	"gin-framework/pkg/auth"
	"gin-framework/pkg/util"
	"gin-framework/types/jwt"
	"gin-framework/types/user"
)

type UserService struct{}

var User = UserService{}

// Login 登录操作
func (s UserService) Login(requestParams user.LoginRequest) (any, error) {
	var userInfo user.User
	if err := global.DB.Table("gin_user").Where("account = ?", requestParams.Account).First(&userInfo).Error; err != nil {
		return userInfo, errors.New("用上不存在")
	}
	// 验证密码
	if !util.VerifyPassword(userInfo.Password, requestParams.Password) {
		return "", errors.New("用户名密码错误")
	}
	jwtUser := jwt.JwtUser{Uuid: userInfo.Uuid}
	jwtToken, err := auth.GenerateJwtToken(global.Cfg.Jwt.Secret, global.Cfg.Jwt.TokenExpire, jwtUser, global.Cfg.Jwt.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	return jwtToken, nil
}
