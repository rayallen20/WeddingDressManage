package model

import (
	"WeddingDressManage/lib/db"
	"fmt"
	"time"
)

// OperationType 操作类型
// createCategoryAndDress:创建新品类礼服
// updateCategory:修改品类信息
// addDress:添加礼服
// updateDress:修改礼服信息
// laundry:送洗
// dailyMaintain:商家日常维护
// preCreateOrder:预创建订单
// createOrder:创建订单
// updateOrder:修改订单
// batchDelivery:批次出件
// allDelivery:全部出件
// putInStorage:礼服归还入库
// itemMaintain:客户破损导致的礼服维护
// applyDiscardDress:礼服申请销库
// applyGiftDress:礼服申请赠与
var OperationType map[string]string = map[string]string{
	"createCategoryAndDress":"createCategoryAndDress",
	"updateCategory":"updateCategory",
	"addDress":"addDress",
	"updateDress":"updateDress",
	"laundry":"laundry",
	"dailyMaintain":"dailyMaintain",
	"preCreateOrder":"preCreateOrder",
	"createOrder":"createOrder",
	"updateOrder":"updateOrder",
	"batchDelivery":"batchDelivery",
	"allDelivery":"allDelivery",
	"putInStorage":"putInStorage",
	"itemMaintain":"itemMaintain",
	"applyDiscardDress":"applyDiscardDress",
	"applyGiftDress":"applyGiftDress",
}

type OperationLog struct {
	Id int
	Kind string `gorm:"column:type"`
	TargetId int
	Data string `gorm:"text"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (o *OperationLog) Save()  {
	err := db.Db.Create(o).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}