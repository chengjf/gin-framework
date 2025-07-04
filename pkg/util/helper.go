package util

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hashicorp/go-uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/schema"
)

// InAnySlice 判断某个字符串是否在字符串切片中
func InAnySlice[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// GenerateBaseSnowId 生成雪花算法ID
func GenerateBaseSnowId(num int, n *snowflake.Node) string {
	if n == nil {
		localIp, err := GetLocalIpToInt()
		if err != nil {
			return ""
		}
		node, _ := snowflake.NewNode(int64(localIp) % 1023)
		n = node
	}
	id := n.Generate()
	switch num {
	case 2:
		return id.Base2()
	case 32:
		return id.Base32()
	case 36:
		return id.Base36()
	case 58:
		return id.Base58()
	case 64:
		return id.Base64()
	default:
		return gconv.String(id.Int64())
	}
}

// GenerateUuid 生成随机字符串
func GenerateUuid(size int) string {
	str, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return gstr.SubStr(str, 0, size)
}

// GeneratePasswordHash 生成密码hash值
func GeneratePasswordHash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

// FormatToString 格式化转化成string
func FormatToString(originStr any) string {
	str := ""
	switch v := originStr.(type) {
	case float64:
		str = strconv.FormatFloat(v, 'f', 10, 64)
	case float32:
		str = strconv.FormatFloat(float64(v), 'f', 10, 32)
	case nil:
		str = ""
	case int:
		str = strconv.FormatInt(int64(v), 10)
	case int32:
		str = strconv.FormatInt(int64(v), 10)
	case int64:
		str = strconv.FormatInt(v, 10)
	default:
		str = fmt.Sprintf("%v", v)
	}
	return str
}

// IsPathExist 判断所给路径文件/文件夹是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// MakeMultiDir 调用os.MkdirAll递归创建文件夹
func MakeMultiDir(filePath string) error {
	if !IsPathExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// MakeFileOrPath 创建文件/文件夹
func MakeFileOrPath(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// WriteContentToFile
// @Description: 写文件
// @param filePath
// @return error
func WriteContentToFile(file *multipart.FileHeader, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	fileBytes, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	if _, err := f.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

// MakeTimeFormatDir
// @Description: 创建时间格式的目录 如：upload/{path}/2023-01-07/
// @param rootPath 根目录
// @param pathName 子目录名称
// @param timeFormat 时间格式 如：2006-01-02、20060102
// @return string
// @return error
func MakeTimeFormatDir(rootPath, pathName, timeFormat string) (string, error) {
	filePath := "upload/"
	if pathName != "" {
		filePath += pathName + "/"
	}
	filePath += time.Now().Format(timeFormat) + "/"
	if err := MakeMultiDir(rootPath + filePath); err != nil {
		return "", err
	}
	return filePath, nil
}

// String2Int 将数组的string转int
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

// GetStructColumnName 获取结构体中的字段名称 _type: 1: 获取tag字段值 2：获取结构体字段值
func GetStructColumnName(s any, _type int) ([]string, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	var fields []string
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 处理匿名字段（嵌套结构体）
		if field.Anonymous {
			// 递归获取嵌套结构体的字段
			embeddedFields, err := GetStructColumnName(v.Field(i).Interface(), _type)
			if err != nil {
				return nil, err
			}
			fields = append(fields, embeddedFields...)
			continue
		}

		var fieldName string
		if _type == 1 {
			fieldName = field.Tag.Get("json")
			if fieldName == "" {
				tagSetting := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
				fieldName = tagSetting["COLUMN"]
			}
		} else {
			fieldName = field.Name
		}

		if fieldName != "" {
			fields = append(fields, fieldName)
		}
	}

	return fields, nil
}

// GetProjectModuleName 获取当前项目的module名称
func GetProjectModuleName() string {
	cmd := exec.Command("go", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(output), "\n")
}
