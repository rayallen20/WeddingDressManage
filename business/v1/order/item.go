package order

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/model"
)

type Item struct {
	Id          int
	Order       *Order
	Dress       *dress.Dress
	MaintainFee int
	Status      string
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
