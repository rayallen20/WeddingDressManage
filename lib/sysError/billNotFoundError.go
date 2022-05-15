package sysError

import "strconv"

type BillNotFoundError struct {
	OrderId int
}

func (b *BillNotFoundError) Error() string {
	return "不存在ID为 " + strconv.Itoa(b.OrderId) + " 的订单所属的押金支付账单"
}
