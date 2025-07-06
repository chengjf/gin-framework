package backend

import (
	"errors"
	"gin-framework/global"
	"time"

	"gin-framework/models"

	"gin-framework/pkg/paginator"
	"gin-framework/pkg/util"

	"gin-framework/types/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Create(requestParams user.UserCreateRequest, ctx *gin.Context) error {

	var (
		user    models.GinUser
		uuid    = util.GenerateBaseSnowId(0, nil)
		pass, _ = util.GeneratePasswordHash(requestParams.Password)
		timeNow = time.Now()
	)
	if err := global.DB.First(&user, "account = ?", requestParams.Account).Error; err == nil {
		return errors.New("账号名已存在")
	}
	localIp, err := util.GetLocalIp()
	if err != nil {
		return err
	}
	user = models.GinUser{
		Uuid:         uuid,
		Account:      requestParams.Account,
		Password:     pass,
		RegisterTime: util.FormatTime(timeNow),
		RegisterIp:   localIp,
		LoginIp:      "",
		Status:       1,
	}
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

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

	builder := paginator.NewBuilder().
		WithDB(global.DB.WithContext(c)).
		WithModel(models.GinUser{}).
		//WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"}).
		//WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"}).
		WithMultiFields(multiFields).
		WithJoins("left", []paginator.OnJoins{{
			LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
			RightTableField: paginator.JoinTableField{Table: models.GinUserProfileTbName, Field: "user_id"},
		}})
	// 查询条件
	if requestParams.Account != "" {
		builder.WithCondition("account like ?", "%"+requestParams.Account+"%")
	}

	pagination, err := builder.Pagination(&userList, requestParams.PageNo, requestParams.PageSize)
	return pagination, err
}

// GetList 获取列表
func (s *UserService) GetList(requestParams user.IndexRequest, c *gin.Context) (any, error) {
	global.Logger.WithContext(c).Info("GetList", requestParams)
	var userList = make([]user.GinUser, 0)
	builder := paginator.NewBuilder().
		WithDB(global.DB.WithContext(c)).
		WithModel(models.GinUser{}).
		WithPreload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		})
	// 查询条件
	if requestParams.Account != "" {
		builder.WithCondition("account like ?", "%"+requestParams.Account+"%")
	}
	if requestParams.Status > 0 {
		builder.WithCondition("status = ?", requestParams.Status)
	}
	pagination, err := builder.Pagination(&userList, requestParams.PageNo, requestParams.PageSize)

	return pagination, err
}
