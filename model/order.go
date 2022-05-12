package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// SaleStrategy 销售策略
// originalPrice:原价
// discount:打折
// customPrice:自定义租金与押金
var SaleStrategy map[string]string = map[string]string{
	"originalPrice": "originalPrice",
	"discount":      "discount",
	"customPrice":   "customPrice",
}

// OrderStatus 订单状态
// unsettled:未付租金
// notYetDelivery:租金已付,未出件
// deliving:出件中
// deliveryFinish:出件结束
// retrieving:还件中
// finish:订单结束
var OrderStatus map[string]string = map[string]string{
	"unsettled":      "unsettled",
	"notYetDelivery": "notYetDelivery",
	"deliving":       "deliving",
	"deliveryFinish": "deliveryFinish",
	"retrieving":     "retrieving",
	"finish":         "finish",
}

type Order struct {
	Id                     int
	CustomerId             int
	Customer               *Customer `gorm:"foreignKey:CustomerId"`
	SerialNumber           string
	Comment                string `gorm:"text"`
	WeddingDate            time.Time
	OriginalCharterMoney   int
	OriginalCashPledge     int
	SaleStrategy           string
	Discount               float64
	DuePayCharterMoney     int
	DuePayCashPledge       int
	DueRefundCashPledge    int
	ActualPayCharterMoney  int
	ActualPayCashPledge    int
	ActualRefundCashPledge int
	TotalMaintainFee       int
	Status                 string
	CreatedTime            time.Time `gorm:"autoCreateTime"`
	UpdatedTime            time.Time `gorm:"autoUpdateTime"`
}

func (o *Order) FindMaxOrderBySNPrefix(date string) error {
	return db.Db.Order("created_time desc").Where("serial_number LIKE ?", date+"%").Limit(1).First(o).Error
}

func (o *Order) Create(rentPlans []*DressRentPlan, items []*OrderItem, billMap map[*Bill][]*BillLog) error {
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

	for _, item := range items {
		item.OrderId = o.Id
	}

	if err := tx.Create(items).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 此处收集item的dressId和itemId
	// 后续写入billLog时需要通过dressId查找该map以确认billLog对应的itemId
	// 后续写入dressRentPlan时需要通过dressId查找该map以确认dressRentPlan对应的itemId
	dressIdMapItemId := make(map[int]int, len(items))
	for _, item := range items {
		dressIdMapItemId[item.DressId] = item.Id
	}

	for bill, billLogs := range billMap {
		bill.OrderId = o.Id
		if err := tx.Create(bill).Error; err != nil {
			tx.Rollback()
			return err
		}

		for _, billLog := range billLogs {
			billLog.BillId = bill.Id
			// 此处需要判断billLog中保存的DressId在item集合中是否存在
			// 因为billLog有可能没有保存DressId 这种billLog也就不需要itemId了
			if itemId, ok := dressIdMapItemId[billLog.DressId]; ok {
				billLog.OrderItemId = itemId
			}
			if err := tx.Create(billLog).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	for _, rentPlan := range rentPlans {
		rentPlan.OrderId = o.Id
		rentPlan.OrderItemId = dressIdMapItemId[rentPlan.DressId]
		if err := tx.Create(rentPlan).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// CountDeliveries 统计订单状态为 "租金已付,未出件" 或 "出件中"的订单数量
func (o *Order) CountDeliveries() (count int64, err error) {
	err = db.Db.Where("status", OrderStatus["notYetDelivery"]).Or("status", OrderStatus["deliving"]).
		Where(o).Find(o).Count(&count).Error
	return count, err
}

// FindDeliveries 分页查找订单状态为 "租金已付,未出件" 或 "出件中"的订单集合
func (o *Order) FindDeliveries(currentPage, itemPerPage int) (orders []*Order, err error) {
	orders = make([]*Order, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Where("status", OrderStatus["notYetDelivery"]).
		Or("status", OrderStatus["deliving"]).Preload("Customer").Order("wedding_date asc").Find(&orders).Error
	return orders, err
}
