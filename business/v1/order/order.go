package order

import (
	"WeddingDressManage/business/v1/customer"
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/business/v1/order/saleStrategy"
	"WeddingDressManage/lib/helper/paramHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	categoryRequest "WeddingDressManage/param/request/v1/category"
	requestParam "WeddingDressManage/param/request/v1/order"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// DateBit 订单号中表示日期部分的位数
const DateBit = 8

// CreateOrderBillNumber 创建订单时所需创建的账单条目数量
// 3条账单分别为 实付租金账单 应付押金账单 应退押金账单
const CreateOrderBillNumber = 3

// CustomPriceMinProportion 自定义租金金额最少占原始租金金额的百分比
const CustomPriceMinProportion = 0.8

// MinDiscount 折扣的最低比例
const MinDiscount = 0.85

// MaxDiscount 折扣的最高比例
const MaxDiscount = 1.0

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

func (o *Order) CalcDiscount(param *requestParam.DiscountParam) error {
	// step1. 校验折扣字段是否符合业务规范
	discount, _ := strconv.ParseFloat(param.Discount, 64)
	if discount < MinDiscount || discount > MaxDiscount {
		return &sysError.DiscountInvalidError{
			Min: MinDiscount,
			Max: MaxDiscount,
		}
	}

	// step2. 将礼服id集合转换为item集合
	dressIds := make([]int, 0, len(param.Dresses))
	for _, dressParam := range param.Dresses {
		dressIds = append(dressIds, dressParam.Id)
	}
	itemBiz := &Item{}
	items, err := itemBiz.Create(dressIds)
	if err != nil {
		return err
	}
	o.Items = items

	// step3. 计算原价
	o.calcOriginalPrice()

	// step4. 计算折扣后价格
	strategy := &saleStrategy.SaleStrategy{
		Type:                 saleStrategy.Discount,
		Discount:             discount,
		OriginalCharterMoney: o.OriginalCharterMoney,
		OriginalCashPledge:   o.OriginalCashPledge,
	}
	dressBizs := make([]*dress.Dress, 0, len(o.Items))
	for _, item := range o.Items {
		dressBiz := item.Dress
		dressBizs = append(dressBizs, dressBiz)
	}
	strategy.Dresses = dressBizs
	err = strategy.CalcPrice()
	if err != nil {
		return err
	}
	o.SaleStrategy = strategy.Type
	o.Discount = strategy.Discount
	o.DuePayCharterMoney = strategy.DuePayCharterMoney
	o.DuePayCashPledge = strategy.DuePayCashPledge
	return nil
}

func (o *Order) Create(param *requestParam.CreateParam) error {
	// step1. 确认客户ID
	customerBiz, err := o.confirmCustomer(param)
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	if customerBiz.Status == model.CustomerStatus["banned"] {
		return &sysError.CustomerBeBannedError{
			Name:   customerBiz.Name,
			Mobile: customerBiz.Mobile,
		}
	}
	o.Customer = customerBiz

	// step2. 校验参数是否符合业务逻辑要求
	err = o.confirmParam(param)
	if err != nil {
		return err
	}

	// step3. 创建item
	err = o.createItem(param)
	if err != nil {
		return err
	}

	// step4. 计算原始租金价格 原始押金价格 婚期
	o.calcOriginalPrice()
	o.WeddingDate = time.Time(param.Order.WeddingDate)

	// step5. 根据销售策略 计算应付租金金额 应付押金金额 应退押金金额
	err = o.calcDuePayPrice(param)
	if err != nil {
		return err
	}
	o.DueRefundCashPledge = o.DuePayCashPledge
	if o.SaleStrategy == saleStrategy.CustomPrice {
		// 自定义租金的金额不得低于原价的80%
		MinProportion := int(float64(o.OriginalCharterMoney) * CustomPriceMinProportion)
		if o.DuePayCharterMoney < MinProportion {
			return &sysError.CustomPriceTooFewError{FloorPrice: MinProportion}
		}
	}

	// step6. 生成订单号
	err = o.genSerialNumber()
	if err != nil {
		return err
	}

	// step7. 确认实付金额
	o.verifyActualPayPrice(param)

	// step8. 创建账目相关信息
	bills := o.genBillForCreate(param.Order.PledgeIsSettled)

	// step9. 创建租赁计划信息
	rentPlans := o.createRentPlan()

	// step10. 落盘至DB
	err = o.save(rentPlans, bills)
	if err != nil {
		return &sysError.DbError{RealError: err}
	}
	return nil
}

// confirmCustomer 确认用户信息
func (o *Order) confirmCustomer(param *requestParam.CreateParam) (customerBiz *customer.Customer, err error) {
	customerBiz = &customer.Customer{
		Name:   param.Customer.Name,
		Mobile: param.Customer.Mobile,
	}
	err = customerBiz.FindOrCreateUser()
	return customerBiz, err
}

// confirmParam 校验参数是否符合业务逻辑要求
// 1. 校验婚期是否早于当天
// 2. 校验折扣是否合规
func (o *Order) confirmParam(param *requestParam.CreateParam) error {
	// 校验婚期是否早于当天
	t := time.Now().AddDate(0, 0, 1)
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if param.Order.WeddingDate.IsBefore(today) {
		return &sysError.DateBeforeTodayError{
			Field: "weddingDate",
		}
	}

	// 校验折扣是否合规
	if param.Order.SaleStrategy.Type == saleStrategy.Discount {
		discount, _ := strconv.ParseFloat(param.Order.SaleStrategy.Discount, 64)
		if discount < MinDiscount || discount > MaxDiscount {
			return &sysError.DiscountInvalidError{
				Min: MinDiscount,
				Max: MaxDiscount,
			}
		}
	}
	return nil
}

// createItem 创建订单操作中 创建item步骤
func (o *Order) createItem(param *requestParam.CreateParam) error {
	dressIds := make([]int, 0, len(param.Order.Items))
	for _, itemParam := range param.Order.Items {
		dressIds = append(dressIds, itemParam.Dress.Id)
	}
	itemBiz := &Item{}
	items, err := itemBiz.Create(dressIds)
	if err != nil {
		return err
	}
	o.Items = items
	return nil
}

// calcOriginalPrice 创建订单操作中 计算原始租金与押金
func (o *Order) calcOriginalPrice() {
	for _, item := range o.Items {
		o.OriginalCharterMoney += item.Dress.Category.CharterMoney
		o.OriginalCashPledge += item.Dress.Category.CashPledge
	}
}

// calcDuePayPrice 创建订单操作中 计算不同的优惠策略后的价格
func (o *Order) calcDuePayPrice(param *requestParam.CreateParam) error {
	// 实际上这三个参数也可以在确认用到时再转换类型 但没有必要 因为后续的策略模式中
	// 只在对应的策略中才会使用特定的参数
	discount, _ := strconv.ParseFloat(param.Order.SaleStrategy.Discount, 64)
	customCharterMoney, _ := paramHelper.ConvertYuanToPenny(param.Order.SaleStrategy.CustomPriceCharter)
	customCashPledge, _ := paramHelper.ConvertYuanToPenny(param.Order.SaleStrategy.CustomPricePledge)

	strategy := &saleStrategy.SaleStrategy{
		Type:                 param.Order.SaleStrategy.Type,
		Discount:             discount,
		OriginalCharterMoney: o.OriginalCharterMoney,
		OriginalCashPledge:   o.OriginalCashPledge,
		CustomCharterMoney:   customCharterMoney,
		CustomCashPledge:     customCashPledge,
	}

	// 提取items中的礼服信息 若优惠策略为打折 则需要根据每一件礼服的价格进行折扣计算
	dressBizs := make([]*dress.Dress, 0, len(o.Items))
	for _, item := range o.Items {
		dressBiz := item.Dress
		dressBizs = append(dressBizs, dressBiz)
	}
	strategy.Dresses = dressBizs

	err := strategy.CalcPrice()
	if err != nil {
		return err
	}
	o.SaleStrategy = strategy.Type
	if o.SaleStrategy == saleStrategy.Discount {
		o.Discount = strategy.Discount
	}
	o.DuePayCharterMoney = strategy.DuePayCharterMoney
	o.DuePayCashPledge = strategy.DuePayCashPledge
	return nil
}

// genSerialNumber 生成订单号
// 订单号规则:订单创建日期 + 3位数字
// 3位数字规则:从001开始 当天每创建一个订单 则数字累加 次日从001开始重新计算
func (o *Order) genSerialNumber() error {
	todayDate := time.Now().Format("2006-01-02")
	todayDateSlice := strings.Split(todayDate, "-")
	snPrefix := ""
	for _, todayDatePart := range todayDateSlice {
		snPrefix += todayDatePart
	}
	orm := &model.Order{}
	err := orm.FindMaxOrderBySNPrefix(snPrefix)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	// 若当天不存在订单 则直接赋001即可
	if errors.Is(err, gorm.ErrRecordNotFound) {
		o.SerialNumber = snPrefix + "001"
		return nil
	}

	// 获取当天订单的最大序号
	maxSNSlice := strings.Split(orm.SerialNumber, "")
	maxSequence := ""
	for i := DateBit; i < len(maxSNSlice); i++ {
		maxSequence += maxSNSlice[i]
	}
	maxSequenceFigure, _ := strconv.Atoi(maxSequence)
	maxSequenceFigure += 1
	maxSequence = fmt.Sprintf("%03d", maxSequenceFigure)
	o.SerialNumber = snPrefix + maxSequence
	return nil
}

// verifyActualPayPrice 计算实付租金与押金
func (o *Order) verifyActualPayPrice(param *requestParam.CreateParam) {
	o.ActualPayCharterMoney = o.DuePayCharterMoney
	if param.Order.PledgeIsSettled {
		o.ActualPayCashPledge = o.DuePayCashPledge
	}
}

func (o *Order) genBillForCreate(isSettled bool) (bills []*Bill) {
	bills = make([]*Bill, 0, CreateOrderBillNumber)

	// 实付租金账单
	collectCharterMoneyBill := &Bill{}
	collectCharterMoneyBill.createCollectCharterMoney(o)
	bills = append(bills, collectCharterMoneyBill)

	// 实付押金账单
	collectCashPledgeBill := &Bill{}
	collectCashPledgeBill.createCollectCashPledge(o, isSettled)
	bills = append(bills, collectCashPledgeBill)

	// 应退押金账单
	restituteCashPledgeBill := &Bill{}
	restituteCashPledgeBill.createRestituteCashPledge(o)
	bills = append(bills, restituteCashPledgeBill)
	return bills
}

// createRentPlan 根据订单信息创建礼服租赁计划
func (o *Order) createRentPlan() []*RentPlan {
	rentPlanBiz := &RentPlan{}
	return rentPlanBiz.create(o)
}

// save 将订单信息与账目信息落盘至DB
func (o *Order) save(rentPlans []*RentPlan, bills []*Bill) error {
	orderORM := &model.Order{
		CustomerId:             o.Customer.Id,
		SerialNumber:           o.SerialNumber,
		WeddingDate:            o.WeddingDate,
		OriginalCharterMoney:   o.OriginalCharterMoney,
		OriginalCashPledge:     o.OriginalCashPledge,
		SaleStrategy:           o.SaleStrategy,
		Discount:               o.Discount,
		DuePayCharterMoney:     o.DuePayCharterMoney,
		DuePayCashPledge:       o.DuePayCashPledge,
		ActualPayCharterMoney:  o.ActualPayCharterMoney,
		ActualPayCashPledge:    o.ActualPayCashPledge,
		ActualRefundCashPledge: o.ActualRefundCashPledge,
		Status:                 model.OrderStatus["notYetDelivery"],
	}

	itemORMs := make([]*model.OrderItem, 0, len(o.Items))
	for _, item := range o.Items {
		itemORM := item.genORMForCreate()
		itemORMs = append(itemORMs, itemORM)
	}

	rentPlanORMs := make([]*model.DressRentPlan, 0, len(rentPlans))
	for _, rentPlan := range rentPlans {
		rentPlanORM := rentPlan.genORMForCreate()
		rentPlanORMs = append(rentPlanORMs, rentPlanORM)
	}

	billORMMap := make(map[*model.Bill][]*model.BillLog, len(bills))
	for _, bill := range bills {
		billORM, billLogORMs := bill.genORMForCreateOrder()
		billORMMap[billORM] = billLogORMs
	}

	return orderORM.Create(rentPlanORMs, itemORMs, billORMMap)
}
