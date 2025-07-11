package global

import (
	"gin-framework/config"
	"gin-framework/pkg/lib"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB      // Mysql数据库
	Logger *lib.Logger   // 日志
	Redis  *redis.Client // redis连接池
	Router *gin.Engine   // 路由
	Cfg    *config.Conf  // yaml配置
)
