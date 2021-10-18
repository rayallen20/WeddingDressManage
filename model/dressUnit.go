package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

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

func (u DressUnit) FindByCategoryIdAndStatus(categoryId, page int, status []string) (usableUnits []DressUnit, err error) {
	usableUnits = make([]DressUnit, 0)
	res := db.Db.Scopes(db.Paginate(page)).Where("category_id = ?", categoryId).
		Where("status IN ?", status).Order("id asc").Find(&usableUnits)
	return usableUnits, res.Error
}

// CountUsableByCategoryId 统计指定礼服品类下 状态为可用(可租/已租/送洗)的礼服数量
func (u DressUnit) CountUsableByCategoryId(categoryId int) (int64, error) {
	var count int64
	res := db.Db.Table("dress_unit").Where("category_id = ?", categoryId).
		Where("status IN ?", []string{UnitStatus["rentable"], UnitStatus["rentOut"], UnitStatus["laundry"]}).
		Count(&count)
	return count, res.Error
}

// CountUnusableByCategoryId 统计指定礼服品类下 状态为不可用(赠与/废弃)的礼服数量
func (u DressUnit) CountUnusableByCategoryId(categoryId int) (int64, error) {
	var count int64
	res := db.Db.Table("dress_unit").Where("category_id = ?", categoryId).
		Where("status IN ?", []string{UnitStatus["gift"], UnitStatus["obsolete"]}).
		Count(&count)
	return count, res.Error
}