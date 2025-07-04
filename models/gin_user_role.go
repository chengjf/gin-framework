package models

import "gin-framework/pkg/util"

var GinUserRoleTbName = "gin_user_role"

type GinUserRole struct {
	Id        uint64          `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	UserId    uint64          `gorm:"column:user_id;type:bigint unsigned;NULL;comment:用户ID" json:"user_id"`      // 用户ID
	RoleId    uint64          `gorm:"column:role_id;type:bigint unsigned;NULL;comment:角色ID 例如：1" json:"role_id"` // 角色ID 例如：1
	CreatedAt util.FormatTime `gorm:"column:created_at;type:datetime;NULL;" json:"created_at"`
}

func (GinUserRole) TableName() string {
	return GinUserRoleTbName
}
