package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var DressStatus = map[string]string{
	// 可租借
	"rentable": "rentable",
	// 租借中
	"rentOut": "rentOut",
	// 送洗中
	"laundry": "laundry",
	// 已赠与
	"gift": "gift",
}

// UnitStatus 礼服状态
var UnitStatus = map[string]string{
	// 可租
	"rentable": "rentable",
	// 已租出
	"rentOut": "rentOut",
	// 送洗
	"laundry": "laundry",
	// 废弃
	"obsolete": "obsolete",
	// 赠与
	"gift": "gift",
}

// UnitSize 礼服尺码
var UnitSize = map[string]string{
	"S":"S",
	"M":"M",
	"F":"F",
	"L":"L",
	"XL":"XL",
	"XXL":"XXL",
	"D":"D",
}

type DressUnit struct {
	// 礼服ID
	Id int

	// 礼服品类ID
	CategoryId int

	// 礼服序号
	SerialNumber int

	// 尺码
	Size string

	// 出租次数
	RentNumber int

	// 送洗次数
	LaundryNumber int

	// 封面图
	CoverImg string

	// 副图
	SecondaryImg string `gorm:"type:text"`

	// 状态
	Status string

	// 排序字段
	Sort int

	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (u *DressUnit) FindMaxSerialNumberByCategoryId() error {
	res := db.Db.Debug().Select("serial_number").Where(u).Order("serial_number desc").First(u)
	return res.Error
}

func (u DressUnit) AddUnitsAndUpdateCategory(units []*DressUnit, category *DressCategory) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新品类信息中的 品类下礼服总数 和 品类可租礼服数 信息
	if err := tx.Table("dress_category").Updates(category).Error; err != nil {
		return err
	}

	// 创建礼服
	if err := tx.Create(units).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}