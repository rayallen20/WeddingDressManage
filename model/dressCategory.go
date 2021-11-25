package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// CategoryStatus 品类状态
// onSale:在售
// soldOut:脱销
// haltSales:下架
var CategoryStatus = map[string]string {
	"onSale":"onSale",
	"soldOut":"soldOut",
	"haltSales":"haltSales",
}

// DressCategory dress_category表的ORM
type DressCategory struct {
	Id int
	KindId int
	Kind *DressKind `gorm:"foreignKey:KindId"`
	SerialNumber string
	Quantity int
	RentableQuantity int
	CharterMoney int
	AvgCharterMoney int
	CashPledge int
	RentCounter int
	LaundryCounter int
	MaintainCounter int
	CoverImg string
	SecondaryImg string `gorm:"text"`
	Status string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

// FindBySerialNumber 根据礼服品类编码 查找1条品类信息 默认为非脱销状态
func (c *DressCategory) FindBySerialNumber() error {
	return db.Db.Not("status = ?", CategoryStatus["haltSales"]).Preload("Kind").Where(c).First(c).Error
}

// AddCategoryAndDresses 使用事务同时创建礼服品类信息和礼服信息
func (c *DressCategory) AddCategoryAndDresses(dressORMs []*Dress) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, dressORM := range dressORMs {
		dressORM.CategoryId = c.Id
	}

	if err := tx.Create(dressORMs).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// FindNormal 查询状态不为haltSales的礼服品类信息
func (c *DressCategory) FindNormal(currentPage, itemPerPage int) (categories []*DressCategory, err error) {
	categories = make([]*DressCategory, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Not("status", CategoryStatus["haltSales"]).Preload("Kind").Find(&categories).Error
	return categories, err
}