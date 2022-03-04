package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// LaundryStatus 礼服送洗记录状态
// underway:送洗中
// finish:送洗完毕
var LaundryStatus map[string]string = map[string]string{
	"underway": "underway",
	"finish":   "finish",
}

type LaundryRecord struct {
	Id                int
	DressId           int
	Dress             *Dress `gorm:"foreignKey:DressId"`
	DirtyPositionImg  string `gorm:"text"`
	Note              string `gorm:"text"`
	StartLaundryDate  time.Time
	DueEndLaundryDate time.Time
	// TODO gorm使用问题:若字段为time.Time(其他类型尚未遇到) 则Save()/Create()时 该字段为空值时 默认该字段值不为空 需在更新时 使用Select()方法指定更新字段
	EndLaundryDate time.Time
	Status         string
	CreatedTime    time.Time `gorm:"autoCreateTime"`
	UpdatedTime    time.Time `gorm:"autoUpdateTime"`
}

// FindUnderway 查询状态为underway的送洗信息
func (l *LaundryRecord) FindUnderway(currentPage, itemPerPage int) (laundryRecords []*LaundryRecord, err error) {
	laundryRecords = make([]*LaundryRecord, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Where("status", LaundryStatus["underway"]).Preload("Dress").Preload("Dress.Category").Preload("Dress.Category.Kind").
		Find(&laundryRecords).Error
	return laundryRecords, err
}

// CountUnderway 统计状态为underway的送洗信息条目
func (l *LaundryRecord) CountUnderway() (count int64, err error) {
	err = db.Db.Where("status", LaundryStatus["underway"]).Find(&LaundryRecord{}).Count(&count).Error
	return count, err
}

func (l *LaundryRecord) FindById() (err error) {
	return db.Db.Where("status", LaundryStatus["underway"]).Preload("Dress").First(l).Error
}

// Finish 送洗结束 送洗记录状态由underway修改为finish 送洗礼服状态由laundry修改为onSale
func (l *LaundryRecord) Finish(dress *Dress) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Updates(l).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Updates(dress).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
