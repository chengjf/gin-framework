package models

var GinUserProfileTbName = "gin_user_profile"

// GinUserProfile 用户表
type GinUserProfile struct {
	Model
	UserId string `gorm:"column:user_id;type:varchar(32);NOT NULL;comment:用户id" json:"user_id"` // 用户id
	Phone  string `gorm:"column:phone" json:"phone"`
}

func (GinUserProfile) TableName() string {
	return GinUserProfileTbName
}
