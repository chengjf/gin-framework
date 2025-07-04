package backend

import (
	"gin-framework/global"

	"gin-framework/models"

	"gin-framework/pkg/paginator"

	"gin-framework/types/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct{}

var User = &UserService{}

// GetIndex 获取列表
func (s *UserService) GetIndex(requestParams user.IndexRequest, c *gin.Context) (any, error) {
	global.Logger.WithContext(c).Info("GetIndex", requestParams)
	x, _ := global.Redis.Get(c, "a").Result()
	global.Logger.WithContext(c).Info("GetIndex", requestParams, x)
	var userList = make([]user.UserList, 0)
	multiFields := []paginator.SelectTableField{
		{Model: models.GinUser{}, Table: models.GinUserTbName, Field: []string{"password", "salt", "_omit"}},
		{Model: models.GinUserProfile{}, Table: models.GinUserProfileTbName, Field: []string{"phone"}},
	}
	pagination, err := paginator.NewBuilder().
		WithDB(global.DB.WithContext(c)).
		WithModel(models.GinUser{}).
		//WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"}).
		//WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"}).
		WithMultiFields(multiFields).
		WithJoins("left", []paginator.OnJoins{{
			LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
			RightTableField: paginator.JoinTableField{Table: models.GinUserProfileTbName, Field: "user_id"},
		}}).
		Pagination(&userList, requestParams.PageNo, requestParams.PageSize)
	return pagination, err
}

// GetList 获取列表
func (s *UserService) GetList(requestParams user.IndexRequest, c *gin.Context) (any, error) {
	var userList = make([]user.GinUser, 0)
	pagination, err := paginator.NewBuilder().
		WithDB(global.DB.WithContext(c)).
		WithModel(models.GinUser{}).
		WithPreload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Pagination(&userList, requestParams.PageNo, requestParams.PageSize)
	return pagination, err
}
