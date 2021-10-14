package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var KindStatus = map[string]string{
	"usable":"usable",
	"unusable":"unusable",
}

// DressKind 礼服品类名称编码表
type DressKind struct {
	// 主键自增ID
	Id int

	// 品类名称
	Name string

	// 品类编码
	Code string

	// 品类状态(可用/不可用)
	Status string

	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

// FindByKindAndCode 根据品类名称和品类编码查询1条品类编码信息
func (k *DressKind) FindByKindAndCode(name, code string) (err error) {
	k.Name = name
	k.Code = code
	k.Status = KindStatus["usable"]
	res := db.Db.Where(k).Find(k)
	return res.Error
}

// FindAllUsableKinds 查询所有可用品类名称和对应的编码信息
func (k DressKind) FindAllUsableKinds() (kinds []DressKind, err error) {
	res := db.Db.Where("status = ?", KindStatus["usable"]).Order("id asc").Find(&kinds)
	return kinds, res.Error
}