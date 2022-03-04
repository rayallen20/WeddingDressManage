package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// OperationType 操作类型
// createCategoryAndDress:创建新品类礼服
// updateCategory:修改品类信息
// addDress:添加礼服
// updateDress:修改礼服信息
// laundry:送洗
// laundryGiveBack:送洗归还
// dailyMaintain:商家日常维护
// dailyMaintainGiveBack:商家日常维护归还
// preCreateOrder:预创建订单
// createOrder:创建订单
// updateOrder:修改订单
// batchDelivery:批次出件
// allDelivery:全部出件
// putInStorage:礼服归还入库
// itemMaintain:客户破损导致的礼服维护
// itemMaintainGiveBack:客户破损导致的礼服维护归还
// applyDiscardDress:礼服申请销库
// applyGiftDress:礼服申请赠与
var OperationType map[string]string = map[string]string{
	"createCategoryAndDress": "createCategoryAndDress",
	"updateCategory":         "updateCategory",
	"addDress":               "addDress",
	"updateDress":            "updateDress",
	"laundry":                "laundry",
	"laundryGiveBack":        "LaundryGiveBack",
	"dailyMaintain":          "dailyMaintain",
	"dailyMaintainGiveBack":  "dailyMaintainGiveBack",
	"preCreateOrder":         "preCreateOrder",
	"createOrder":            "createOrder",
	"updateOrder":            "updateOrder",
	"batchDelivery":          "batchDelivery",
	"allDelivery":            "allDelivery",
	"putInStorage":           "putInStorage",
	"itemMaintain":           "itemMaintain",
	"itemMaintainGiveBack":   "itemMaintainGiveBack",
	"applyDiscardDress":      "applyDiscardDress",
	"applyGiftDress":         "applyGiftDress",
}

type OperationLog struct {
	Id          int
	Kind        string `gorm:"column:type"`
	TargetId    int
	Data        string    `gorm:"text"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (o *OperationLog) Save() {
	db.Db.Create(o)
}

func (o *OperationLog) SaveWithSecondaryLog(secondaryLogs []*OperationSecondaryEntity) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(o).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, secondaryLog := range secondaryLogs {
		secondaryLog.OperationLogId = o.Id
	}

	if err := tx.Create(secondaryLogs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
