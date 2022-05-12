package sysError

import "strconv"

type DeliveryOrderNotExist struct {
	Id int
}

func (d *DeliveryOrderNotExist) Error() string {
	return "不存在Id为 " + strconv.Itoa(d.Id) + " 的待出件订单"
}
