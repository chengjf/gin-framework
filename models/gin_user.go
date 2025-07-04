package models

import "gin-framework/pkg/util"

var GinUserTbName = "gin_user"

// GinUser 用户表
type GinUser struct {
	Model
	Uuid         string          `gorm:"column:uuid;type:varchar(32);NOT NULL;comment:唯一id号" json:"uuid"`                           // 唯一id号
	Account      string          `gorm:"column:account;type:varchar(64);NOT NULL;comment:用户名" json:"user_name"`                     // 用户名
	Password     string          `gorm:"column:password;type:varchar(64);NOT NULL;comment:密码" json:"password"`                      // 密码
	Status       uint8           `gorm:"column:status;type:tinyint unsigned;default:1;NOT NULL;comment:状态 1：正常 2：禁用" json:"status"` // 状态 1：正常 2：禁用
	RegisterTime util.FormatTime `gorm:"column:register_time;type:timestamp;NOT NULL;comment:注册时间" json:"register_time"`            // 注册时间
	RegisterIp   string          `gorm:"column:register_ip;type:varchar(32);NOT NULL;comment:注册ip" json:"register_ip"`              // 注册ip
	LoginTime    util.FormatTime `gorm:"column:login_time;type:timestampDEFAULT NULL;comment:登录时间" json:"login_time"`               // 登录时间
	LoginIp      string          `gorm:"column:login_ip;type:varchar(32);DEFAULT NULL;comment:登录ip" json:"login_ip"`                // 登录ip
}

func (GinUser) TableName() string {
	return GinUserTbName
}
