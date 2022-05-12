package order

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/model"
	"strings"
)

// RentStatus 出租状态
// notYetDelivery: 尚未出件
// deliveryFinish: 已出件
var RentStatus map[string]string = map[string]string{
	"notYetDelivery": "notYetDelivery",
	"deliveryFinish": "deliveryFinish",
}

type Item struct {
	Id          int
	Order       *Order
	Dress       *dress.Dress
	MaintainFee int
	Status      string
	RentStatus  string
}

func (i *Item) PreCreate(dressIds []int) ([]*Item, error) {
	dressBiz := &dress.Dress{}
	dresses, err := dressBiz.FindByIds(dressIds)
	if err != nil {
		return nil, err
	}
	items := make([]*Item, 0, len(dressIds))
	for _, dressObj := range dresses {
		item := &Item{
			Dress:       dressObj,
			MaintainFee: 0,
			Status:      model.OrderItemStatus["invalid"],
		}
		items = append(items, item)
	}
	return items, nil
}

func (i *Item) Create(dressIds []int) ([]*Item, error) {
	dressBiz := &dress.Dress{}
	dresses, err := dressBiz.FindByIds(dressIds)
	if err != nil {
		return nil, err
	}
	items := make([]*Item, 0, len(dressIds))
	for _, dressObj := range dresses {
		item := &Item{
			Dress:       dressObj,
			MaintainFee: 0,
			Status:      model.OrderItemStatus["valid"],
		}
		items = append(items, item)
	}
	return items, nil
}

func (i *Item) genORMForCreate() *model.OrderItem {
	return &model.OrderItem{
		DressId:     i.Dress.Id,
		MaintainFee: 0,
		Status:      i.Status,
	}
}

func (i *Item) fill(orm *model.OrderItem, order *Order) {
	i.Id = orm.Id
	i.Order = order
	i.MaintainFee = orm.MaintainFee
	i.Status = orm.Status
	if orm.Dress != nil {
		i.Dress = &dress.Dress{
			Id:         orm.Dress.Id,
			CategoryId: orm.Dress.CategoryId,
			Category: &dress.Category{
				Id: orm.Dress.Category.Id,
				Kind: &dress.Kind{
					Id:     orm.Dress.Category.Kind.Id,
					Name:   orm.Dress.Category.Kind.Name,
					Code:   orm.Dress.Category.Kind.Code,
					Status: orm.Dress.Category.Kind.Status,
				},
				SerialNumber:     orm.Dress.Category.SerialNumber,
				Quantity:         orm.Dress.Category.Quantity,
				RentableQuantity: orm.Dress.Category.RentableQuantity,
				CharterMoney:     orm.Dress.Category.CharterMoney,
				AvgCharterMoney:  orm.Dress.Category.AvgCharterMoney,
				CashPledge:       orm.Dress.Category.CashPledge,
				RentCounter:      orm.Dress.Category.RentCounter,
				LaundryCounter:   orm.Dress.Category.LaundryCounter,
				MaintainCounter:  orm.Dress.Category.MaintainCounter,
				CoverImg:         urlHelper.GenFullImgWebSite(orm.Dress.Category.CoverImg),
				Status:           orm.Dress.Category.Status,
			},
			SerialNumber:    orm.Dress.SerialNumber,
			Size:            orm.Dress.Size,
			RentCounter:     orm.Dress.RentCounter,
			LaundryCounter:  orm.Dress.LaundryCounter,
			MaintainCounter: orm.Dress.MaintainCounter,
			CoverImg:        urlHelper.GenFullImgWebSite(orm.Dress.CoverImg),
			Status:          orm.Dress.Status,
		}

		// 填充品类副图
		if orm.Dress.Category.SecondaryImg != "" {
			i.Dress.Category.SecondaryImg = urlHelper.GenFullImgWebSites(strings.Split(orm.Dress.Category.SecondaryImg, "|"))
		} else {
			i.Dress.Category.SecondaryImg = make([]string, 0)
		}

		// 填充礼服副图
		if orm.Dress.SecondaryImg != "" {
			i.Dress.SecondaryImg = urlHelper.GenFullImgWebSites(strings.Split(orm.Dress.SecondaryImg, "|"))
		} else {
			i.Dress.SecondaryImg = make([]string, 0)
		}

		// 判断租赁状态 实际上租赁状态就是礼服状态 只是为了给controller层方便展示
		// 此处做了一个转换
		if orm.Dress.Status == model.DressStatus["rentOut"] {
			i.RentStatus = RentStatus["deliveryFinish"]
		}

		if orm.Dress.Status == model.DressStatus["onSale"] || orm.Dress.Status == model.DressStatus["preRent"] {
			i.RentStatus = RentStatus["notYetDelivery"]
		}
	}
}
