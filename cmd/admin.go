package cmd

import (
	"fmt"
	"time"

	"gin-framework/config"

	"gin-framework/global"

	"gin-framework/models"

	"gin-framework/pkg/util"

	"gin-framework/bootstrap"

	"github.com/urfave/cli/v2"
)

var (
	account  string
	password string
	roleId   string
)

// AccountCmd 管理者账号创建命令
func AccountCmd() *cli.Command {
	return &cli.Command{
		Name:  "account",
		Usage: "Create a new manager account",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "account",
				Aliases:     []string{"a"},
				Value:       "",
				Usage:       "请输入账号名称 如：admin",
				Destination: &account,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "请输入账号密码 如：admin888",
				Destination: &password,
				Required:    true,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.MysqlService)
			return generateAdmin()
		},
	}
}

// generateAdmin 生成admin信息工具
func generateAdmin() error {
	var (
		user    models.GinUser
		uuid    = util.GenerateBaseSnowId(0, nil)
		pass, _ = util.GeneratePasswordHash(password)
		timeNow = time.Now()
	)
	if err := global.DB.First(&user, "account = ?", account).Error; err == nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "account: "+account+" is already existed")
		return err
	}
	localIp, err := util.GetLocalIp()
	if err != nil {
		return err
	}
	user = models.GinUser{
		Uuid:         uuid,
		Account:      account,
		Password:     pass,
		RegisterTime: util.FormatTime(timeNow),
		RegisterIp:   localIp,
		LoginIp:      "",
		Status:       1,
	}
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	fmt.Printf("\u001B[34m账号：%s 密码：%s 角色ID：%s 生成成功\u001B[0m\n", account, password, roleId)
	return nil
}
