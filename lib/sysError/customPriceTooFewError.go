package sysError

import (
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
)

type CustomPriceTooFewError struct {
	FloorPrice int
}

func (c *CustomPriceTooFewError) Error() string {
	floorPricePennyStr := strconv.Itoa(c.FloorPrice)
	floorPriceYuan := paramHelper.ConvertPennyToYuan(floorPricePennyStr)
	return "custom price charter money must be great than " + floorPriceYuan
}
