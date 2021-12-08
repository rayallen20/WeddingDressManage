package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// GiftRecordStatus 礼服赠与申请状态
var GiftRecordStatus map[string]string = map[string]string{
	"pendingApproval": "pendingApproval",
	"approved":        "approved",
	"refused":         "refused",
}

type DressGiftRecord struct {
	Id              int
	DressCategoryId int
	Category        *DressCategory `gorm:"foreignKey:DressCategoryId"`
	DressId         int
	Dress           *Dress `gorm:"foreignKey:DressId"`
	CustomerId      int
	Customer        *Customer `gorm:"foreignKey:CustomerId"`
	Note            string    `gorm:"text"`
	Status          string
	CreatedTime     time.Time `gorm:"autoCreateTime"`
	UpdatedTime     time.Time `gorm:"autoUpdateTime"`
}

func (g *DressGiftRecord) Save() error {
	return db.Db.Save(g).Error
}
