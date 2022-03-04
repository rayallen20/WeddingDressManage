package sysError

import "strconv"

type DressIsNotLaunderingError struct {
	NotLaunderingId int
}

func (d *DressIsNotLaunderingError) Error() string {
	if d.NotLaunderingId != 0 {
		return "there doesn't exist laundering dress which id = " + strconv.Itoa(d.NotLaunderingId)
	}
	return "laundering dress doesn't exist"
}
