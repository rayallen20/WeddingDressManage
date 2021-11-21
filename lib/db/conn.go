package db

import (
	"WeddingDressManage/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Db 数据库连接句柄
var Db *gorm.DB

// init 初始化数据库连接句柄
func init() {
	// 若句柄不为空 说明已经连接数据库 不需再连接
	if Db != nil {
		return
	}

	dsn := fillConnArgs()
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy {
			// 全局禁用表名复数
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Conn DB failed, err: " + err.Error())
	}
}

// fillConnArgs 根据配置拼接连接数据库的必要信息
func fillConnArgs() (args string) {
	return conf.Conf.Database.User + ":" + conf.Conf.Database.Password +"@tcp(" + conf.Conf.Database.Domain +
		":" + conf.Conf.Database.Port + ")/" + conf.Conf.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
}