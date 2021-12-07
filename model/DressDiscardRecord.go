package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// DiscardRecordStatus 礼服销库申请状态
var DiscardRecordStatus map[string]string = map[string]string{
	"pendingApproval": "pendingApproval",
	"approved":        "approved",
	"refused":         "refused",
}

type DressDiscardRecord struct {
	Id              int
	DressCategoryId int
	Category        *DressCategory `gorm:"foreignKey:DressCategoryId"`
	DressId         int
	Dress           *Dress `gorm:"DressId"`
	Note            string `gorm:"text"`
	Status          string
	CreatedTime     time.Time `gorm:"autoCreateTime"`
	UpdatedTime     time.Time `gorm:"autoUpdateTime"`
}

func (r *DressDiscardRecord) Save() error {
	return db.Db.Save(r).Error
}
