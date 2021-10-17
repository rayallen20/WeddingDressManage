package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var CategoryStatus = map[string]string{
	"usable": "usable",
	"unusable": "unusable",
}

type DressCategory struct {
	// 礼服品类ID
	Id int

	// 礼服类别ID
	KindId int

	// 礼服类别编码
	Code string

	// 礼服编号
	SerialNumber string

	// 可租礼服数量
	RentableQuantity int

	// 该品类礼服的总数量
	Quantity int

	// 租金
	CharterMoney int

	// 押金
	CashPledge int

	// 该品类礼服总共被出租的次数
	RentNumber int

	// 该品类礼服总送洗次数
	LaundryNumber int

	// 平均租金
	AvgRentMoney int

	// 封面图
	CoverImg string

	// 副图
	// Tips:DB中的TEXT类型要在ORM中指明
	// https://stackoverflow.com/questions/64035165/unsupported-data-type-error-on-gorm-field-where-custom-valuer-returns-nil
	SecondaryImg string `gorm:"type:text"`

	// 状态
	Status string

	// 排序字段
	Sort int

	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

// FindByKindIdAndCodeAndSN 根据KindId字段值 Code字段值 和 SerialCode字段值 查找1条信息
func (c *DressCategory) FindByKindIdAndCodeAndSN(kindId int, code, serialNumber string) (err error) {
	c.KindId = kindId
	c.Code = code
	c.SerialNumber = serialNumber
	res := db.Db.Where(c).Find(c)
	return res.Error
}

// AddCategoryAndUnits 使用事务同时添加礼服品类和具体的礼服
func (c *DressCategory) AddCategoryAndUnits(units []*DressUnit) ([]*DressUnit, error) {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, unit := range units {
		unit.CategoryId = c.Id
	}

	if err := tx.Create(units).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return units, tx.Commit().Error
}

func (c DressCategory) FindByStatus(page int) ([]DressCategory, error) {
	var categoryInfos []DressCategory
	res := db.Db.Scopes(db.Paginate(page)).Where("status", CategoryStatus["usable"]).Find(&categoryInfos).Order("id asc")
	return categoryInfos, res.Error
}

// CountUsable 统计状态为可用的礼服品类信息数量
func (c DressCategory) CountUsable() (int64, error) {
	var count int64
	res := db.Db.Table("dress_category").Where("status", CategoryStatus["usable"]).Count(&count)
	return count, res.Error
}

// FindById 根据id字段值查找1条数据
func (c *DressCategory) FindById() error {
	res := db.Db.Where(c).Find(c)
	return res.Error
}

func (c *DressCategory) Update() error {
	res := db.Db.Model(c).Updates(c)
	return res.Error
}