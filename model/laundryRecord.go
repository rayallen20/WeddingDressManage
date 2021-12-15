package model

import (
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
