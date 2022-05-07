package sysError

import "strconv"

type DiscountInvalidError struct {
	Min float64
	Max float64
}

func (d *DiscountInvalidError) Error() string {
	return "discount must between " +
		strconv.FormatFloat(d.Min, 'f', 2, 64) + " to " +
		strconv.FormatFloat(d.Max, 'f', 2, 64)
}
