package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// KindStatus 礼服大类状态
// onSale:在售
// haltSales:下架
var KindStatus = map[string]string{
	"onSale":"onSale",
	"haltSales":"haltSales",
}

// DressKind dress_kind表的ORM
type DressKind struct {
	Id int
	Name string
	Code string
	Status string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

// FindById 根据id字段 查找1条礼服大类信息 默认礼服大类状态为在售
func (k *DressKind) FindById() error {
	res := db.Db.Where("status = ?", KindStatus["onSale"]).First(k)
	return res.Error
}

// FindAllOnSale 查询所有在售状态的大类信息
func (k DressKind) FindAllOnSale() ([]*DressKind, error) {
	orms := make([]*DressKind, 0)
	err := db.Db.Where("status", KindStatus["onSale"]).Find(&orms).Error
	return orms, err
}