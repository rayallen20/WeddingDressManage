package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// MaintainSource 维护记录来源
// item 订单内礼服归还时需要维护(此来源的维护信息需包含订单相关信息 会涉及到钱)
// daily 商家日常维护(此来源的维护信息无订单相关信息 不涉及到钱)
var MaintainSource map[string]string = map[string]string{
	"item":  "item",
	"daily": "daily",
}

// MaintainStatus 维护状态
// underway 维护中
// finish 维护结束
var MaintainStatus map[string]string = map[string]string{
	"underway": "underway",
	"finish":   "finish",
}

type MaintainRecord struct {
	Id     int
	Source string
	// TODO:此处差order的orm 但目前暂时没做到订单模块 故未写
	OrderId int
	// TODO:此处差order_item的orm 但目前暂时没做到订单模块 故未写
	OrderItemId            int
	DressId                int
	Dress                  *Dress `gorm:"foreignKey:DressId"`
	DuePayMaintainMoney    int
	ActualPayMaintainMoney int
	MaintainPositionImg    string `gorm:"text"`
	Note                   string `gorm:"text"`
	StartMaintainDate      time.Time
	PlanEndMaintainDate    time.Time
	EndMaintainDate        time.Time
	// TODO gorm使用问题:若字段为time.Time(其他类型尚未遇到) 则Save()/Create()时 该字段为空值时 默认该字段值不为空 需在更新时 使用Select()方法指定更新字段
	Status      string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (m *MaintainRecord) FindUnderway(currentPage, itemPerPage int) (maintainRecords []*MaintainRecord, err error) {
	maintainRecords = make([]*MaintainRecord, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Where("status", MaintainStatus["underway"]).
		Preload("Dress").Preload("Dress.Category").Preload("Dress.Category.Kind").Find(&maintainRecords).Error
	return maintainRecords, err
}

func (m *MaintainRecord) CountUnderway() (count int64, err error) {
	err = db.Db.Where("status", MaintainStatus["underway"]).Find(m).Count(&count).Error
	return count, err
}

func (m *MaintainRecord) FindById() (err error) {
	return db.Db.Where("status", LaundryStatus["underway"]).Preload("Dress").
		Preload("Dress.Category").First(m).Error
}

// Finish 维护结束 维护记录状态由underway修改为finish 维护礼服状态由maintain修改为onSale
// 维护礼服所属品类的可租赁件数 +1
func (m *MaintainRecord) Finish(dress *Dress, category *DressCategory) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Updates(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Updates(dress).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Updates(category).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}