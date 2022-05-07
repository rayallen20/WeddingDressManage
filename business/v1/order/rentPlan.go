package order

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/model"
	"time"
)

// OnSaleToPreRentDays 礼服进入预租状态的开始日期距离婚期的天数 负数表示婚期之前的天数
const OnSaleToPreRentDays = -3

// PreOnSaleToOnSaleDays 礼服从预上架状态变更为可租赁状态的日期距离婚期的天数 正数表示婚期之后的天数
const PreOnSaleToOnSaleDays = 3

type RentPlan struct {
	Id               int
	Order            *Order
	Item             *Item
	Dress            *dress.Dress
	PreRentStartDate time.Time
	WeddingDate      time.Time
	PreOnSaleEndDate time.Time
	Status           string
}

func (r *RentPlan) create(order *Order) []*RentPlan {
	rentPlans := make([]*RentPlan, 0, len(order.Items))
	for _, item := range order.Items {
		rentPlan := &RentPlan{
			Order:            order,
			Item:             item,
			Dress:            item.Dress,
			PreRentStartDate: order.WeddingDate.AddDate(0, 0, OnSaleToPreRentDays),
			WeddingDate:      order.WeddingDate,
			PreOnSaleEndDate: order.WeddingDate.AddDate(0, 0, PreOnSaleToOnSaleDays),
			Status:           model.DressRentPlanStatus["valid"],
		}
		rentPlans = append(rentPlans, rentPlan)
	}
	return rentPlans
}

func (r *RentPlan) genORMForCreate() (orm *model.DressRentPlan) {
	orm = &model.DressRentPlan{
		OrderSerialNumber: r.Order.SerialNumber,
		DressId:           r.Dress.Id,
		PreRentStartDate:  r.PreRentStartDate,
		WeddingDate:       r.WeddingDate,
		PreOnSaleEndDate:  r.PreOnSaleEndDate,
		Status:            r.Status,
	}
	return orm
}
