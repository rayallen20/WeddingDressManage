package saleStrategy

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
	"strings"
)

type discountStrategy struct {
	discount             float64
	originalCharterMoney int
	// beforeTruncCharterMoney 取整前折扣后的租金
	beforeTruncCharterMoney int
	originalCashPledge      int
	duePayCharterMoney      int
	duePayCashPledge        int
	dresses                 []*dress.Dress
}

func (d *discountStrategy) calcPrice(strategy *SaleStrategy) {
	d.originalCharterMoney = strategy.OriginalCharterMoney
	d.originalCashPledge = strategy.OriginalCashPledge

	d.calcCharterMoney()
	strategy.DuePayCharterMoney = d.duePayCharterMoney

	d.calcCashPledge()
	strategy.DuePayCashPledge = d.duePayCashPledge
}

// calcCharterMoney 折扣策略下计算租金
// 规则:
// step1. 计算订单中每一件礼服的折扣后金额
// step2. 将折扣后的金额精确到分
// (有可能折扣后的金额出现分无法描述的情况 比如原租金 100.5元 折扣比例75折 则折扣后的价格为75.375元 此时最后一位的5就是无法用分描述的情况)
// step3. 累加折扣后的金额
// step4. 然后不要角和分的部分 元的部分个位数取0
// 即为该笔订单折扣后的应付租金
func (d *discountStrategy) calcCharterMoney() {
	for _, dressBiz := range d.dresses {
		// 每件礼服打折后的租金
		afterDiscount := float64(dressBiz.Category.CharterMoney) * d.discount
		// 截取掉打折后租金中 无法用分描述的部分
		truncDiscount := int(afterDiscount)
		d.beforeTruncCharterMoney += truncDiscount
	}
	d.truncCharterMoney()
}

// calcCashPledge 折扣策略下计算押金
// 规则:押金不变
func (d *discountStrategy) calcCashPledge() {
	d.duePayCashPledge = d.originalCashPledge
}

// truncCharterMoney 将折扣优惠后的租金 去掉角和分 元的部分个位数取0
// e.g.
// 折扣优惠后的租金为:223.52元
// 截取后为:220元
func (d *discountStrategy) truncCharterMoney() {
	beforeTruncCharterMoneyYuan := int(float64(d.beforeTruncCharterMoney) * paramHelper.PennyToYuan)
	yuanStr := strconv.Itoa(beforeTruncCharterMoneyYuan)
	yuanSlice := strings.Split(yuanStr, "")
	truncYuanSlice := make([]string, 0, len(yuanSlice))
	// 取元的部分中 除个位数以外的数字
	for i := 0; i <= len(yuanSlice)-2; i++ {
		truncYuanSlice = append(truncYuanSlice, yuanSlice[i])
	}
	// 最后补0
	truncYuanSlice = append(truncYuanSlice, "0")
	truncYuanStr := ""
	for _, yuanFigure := range truncYuanSlice {
		truncYuanStr += yuanFigure
	}
	d.duePayCharterMoney, _ = strconv.Atoi(truncYuanStr)
	d.duePayCharterMoney *= paramHelper.YuanToPenny
}
