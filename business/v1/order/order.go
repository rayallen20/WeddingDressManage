package order

import (
	"WeddingDressManage/business/v1/customer"
	"WeddingDressManage/business/v1/dress"
	categoryRequest "WeddingDressManage/param/request/v1/category"
	requestParam "WeddingDressManage/param/request/v1/order"
	"time"
)

type Order struct {
	Id                     int
	Customer               *customer.Customer
	SerialNumber           string
	WeddingDate            time.Time
	Items                  []*Item
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
}

func (o *Order) Search(param *requestParam.SearchParam) (categories []*dress.Category, totalPage int, totalItem int64, err error) {
	// TODO: 此处由于目前未做筛选 故模拟一个查看全部可用品类请求即可
	var curlParam *categoryRequest.ShowParam = &categoryRequest.ShowParam{Pagination: param.Pagination}
	categoryBiz := &dress.Category{}
	return categoryBiz.Show(curlParam)
}

func (o *Order) PreCreate(param *requestParam.PreCreateParam) error {
	// step1. 将请求中的dressId填充为item
	dressIds := make([]int, 0, len(param.Dresses))
	for _, dressParam := range param.Dresses {
		dressIds = append(dressIds, dressParam.Id)
	}
	itemBiz := &Item{}
	items, err := itemBiz.PreCreate(dressIds)
	if err != nil {
		return err
	}
	o.Items = items
	// step2. 计算原始租金价格 原始押金价格 婚期
	for _, item := range items {
		o.OriginalCharterMoney += item.Dress.Category.CharterMoney
		o.OriginalCashPledge += item.Dress.Category.CashPledge
	}
	o.WeddingDate = time.Time(param.Order.WeddingDate)
	return nil
}
