package db

import (
	"WeddingDressManage/conf"
	"gorm.io/gorm"
)

// Paginate 分页器场景
func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageSize := conf.Conf.DataBase.PageSize
		offset := (page - 1) * pageSize
		return db.Limit(pageSize).Offset(offset)
	}
}
