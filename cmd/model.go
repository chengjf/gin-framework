package cmd

import (
	"fmt"
	"gin-framework/bootstrap"
	"gin-framework/config"
	"gin-framework/global"

	gomodel "github.com/MQEnergy/gorm-model"

	"github.com/urfave/cli/v2"
)

var (
	tbName  string
	mDir    string // 模型存储目录
	prefix  string // 数据表前缀
	dsn     string // 数据库连接信息
	isCover bool   // 是否覆盖
)

// ModelCmd 数据库模型创建命令
func ModelCmd() *cli.Command {
	return &cli.Command{
		Name:  "model",
		Usage: "Create a new model class",
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
				Name:        "tbname",
				Aliases:     []string{"tb"},
				Value:       "all",
				Usage:       "表名称 如：初始化所有（all）单个数据表就填写表名（如：gin_admin）",
				Destination: &tbName,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "mdir",
				Aliases:     []string{"dir"},
				Value:       "./models",
				Usage:       "模型存储目录 如：./models（存入在当前执行命令所在目录，支持多级目录）",
				Destination: &mDir,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "prefix",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "需要过滤的数据表前缀 如: gin_",
				Destination: &prefix,
				Required:    false,
			},
			&cli.BoolFlag{
				Name:        "isCover",
				Aliases:     []string{"ic"},
				Value:       false,
				Usage:       "是否覆盖 true:覆盖已经生成的 false:不覆盖",
				Destination: &isCover,
				Required:    false,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.MysqlService)
			return generateModel()
		},
	}
}

// generateModel 创建model
func generateModel() error {
	modelConfig := &gomodel.ModelConfig{
		DB:      global.DB,
		DbName:  global.Cfg.Mysql[0].DbName,
		MDir:    mDir,
		Prefix:  prefix,
		IsCover: isCover,
	}
	if tbName == "all" {
		strs, errs := gomodel.GenerateAllModel(modelConfig)
		for i, str := range strs {
			if errs[i] != nil {
				fmt.Printf("\x1b[31m%s\x1b[0m\n", str)
			} else {
				fmt.Printf(fmt.Sprintf("\u001B[34m%s\u001B[0m\n", str))
			}
		}
	} else {
		str, err := gomodel.GenerateSingleModel(modelConfig, tbName)
		if err != nil {
			fmt.Printf(fmt.Sprintf("\x1b[31m%s\x1b[0m\n", str))
			return nil
		}
		fmt.Printf(fmt.Sprintf("\u001B[34m%s\u001B[0m\n", str))
	}
	return nil
}
