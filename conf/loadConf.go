package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
)

type Config struct {
	DataBase DataBase
	File File
}

type DataBase struct {
	Domain string
	Port string
	UserName string
	Password string
	DataBaseName string
}

type File struct {
	Path string
	RandNumLen int
	ImgType []string
}

// 本变量用于区分开发/生产环境 默认为开发环境
var env string = "dev"

// 本方法用于读取配置
func (c *Config) load() error {
	checkEnv()
	filePath := GetConfigFilePath()
	fileName := "config." + env + ".yaml"

	v := viper.New()
	v.AddConfigPath(filePath)
	v.SetConfigName(fileName)
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

// 本函数用于确认开发/生产环境
func checkEnv() {
	osEnv := os.Getenv("WEDDING_DRESS_MANAGE_ENV")
	// TODO: 此处由于只有dev环境和prod环境,没有test环境,所以不需要再判断其他可能性
	if osEnv == "prod" {
		env = "prod"
	} else {
		env = "dev"
	}
}

func GetConfigFilePath() string {
	nowFilePath := currentFile()
	filePathSlice := strings.Split(nowFilePath, "/")
	var filePath string
	for i := 0; i < len(filePathSlice) - 2; i++ {
		filePath += filePathSlice[i]
		filePath += "/"
	}
	filePath += "/"
	return filePath
}

func currentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Can not get current file info")
	}
	return file
}

var Conf *Config = &Config{}

func init() {
	err := Conf.load()
	if err != nil {
		panic(fmt.Errorf("load config failed, err:" + err.Error()))
	}
}