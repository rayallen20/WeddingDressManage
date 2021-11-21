package conf

import (
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
)

var Conf *Config = &Config{}

func init()  {
	err := Conf.load()
	if err != nil {
		panic("load config failed:" + err.Error())
	}
}

// 	Config 配置对象
type Config struct {
	Database Database
	Upload Upload
}

// load 读取配置文件至Config对象
func (c *Config) load() error {
	checkEnv()
	configFilePath := getConfigFilePath()
	configFileName := "config." + env

	v := viper.New()
	v.AddConfigPath(configFilePath)
	v.SetConfigName(configFileName)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(c)
	if err != nil {
		return err
	}

	return nil
}

// Database 数据库相关配置
type Database struct {
	// Domain 数据库服务器IP地址
	Domain string
	// Port 数据库端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// Name 数据库名
	Name string
}

// Upload 上传文件相关配置
type Upload struct {
	// Path 文件保存路径
	Path string
}

// env 用于区分生产/研发环境的变量 默认为研发环境
// TODO:当生产环境和研发环境处在同一服务器时 如何使用系统变量区分
// 目的:生产环境:读config.yaml 研发环境:读config.dev.yaml
// 思路1:生产环境没有config.dev.yaml 研发环境有config.dev.yaml	隐藏含义:代码有区分环境的能力
// 思路2:生产和测试环境都用config.yaml 运维人员根据config.template.yaml自行修改	隐藏含义:代码没有区分环境的能力
// 我更倾向于 让代码有区分环境的能力
var env = "dev"

// checkEnv 检查当前环境为生产环境还是研发环境
func checkEnv() {
	envVariable := os.Getenv("WEDDING_DRESS_MANAGE_ENV")
	if envVariable == "prod" {
		env = "prod"
	}
}

// getConfigFilePath 获取配置文件所在路径
func getConfigFilePath() string {
	currentFilePath := getCurrentFilePath()
	currentFilePathSlice := strings.Split(currentFilePath, "/")
	var configFilePath string
	for i := 0; i < len(currentFilePathSlice) - 2; i++ {
		configFilePath += currentFilePathSlice[i]
		configFilePath += "/"
	}
	configFilePath += "/"
	return configFilePath
}

// getCurrentFilePath 获取当前文件所在路径
func getCurrentFilePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Can not get current file info!")
	}
	return file
}

