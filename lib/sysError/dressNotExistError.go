package sysError

import "strconv"

type DressNotExistError struct {
	Id int
}

func (d *DressNotExistError) Error() string {
	if d.Id != 0 {
		return "the dress which id = " + strconv.Itoa(d.Id) + " doesn't exist"
	}
	return "dress doesn't exist"
}
