package user

import (
	"gin-framework/models"
	"gin-framework/pkg/util"
)

type BaseUser models.GinUser
type GinUserInfo models.GinUserRole
type GinUserProfile models.GinUserProfile

type User struct {
	BaseUser
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}

// IndexRequest 获取用户列表请求参数
type IndexRequest struct {
	PageNo   int    `form:"page_no" json:"page_no"`
	PageSize int    `form:"page_size" json:"page_size"`
	Account  string `form:"account" json:"account"`
	Status   int    `form:"status" json:"status"`
}

type UserCreateRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserList joins获取关联列表
type UserList struct {
	BaseUser
	Phone string `json:"phone"`
}

type GinRole struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type ListUser struct {
	Id           uint64          `json:"id"`
	Uuid         string          `json:"uuid"`
	Account      string          `json:"account"`
	Status       uint8           `json:"status"`
	RegisterTime util.FormatTime `json:"register_time"`
	RegisterIp   string          `json:"register_ip"`
	LoginTime    util.FormatTime `json:"login_time"`
	LoginIp      string          `json:"login_ip"`
}

// GinUser preload获取关联列表
type GinUser struct {
	ListUser
	Roles []GinRole `gorm:"many2many:gin_user_role;foreignKey:Id;joinForeignKey:UserId;references:Id;joinReferences:RoleId" json:"roles"`
}

// LoginRequest 用户登录请求参数
type LoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
