package db

import "gorm.io/gorm"

// Paginate 分页器场景
// page:当前页数
// itemPerPage:每页显示条数
func Paginate(page, itemPerPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * itemPerPage
		return db.Limit(itemPerPage).Offset(offset)
	}
}
