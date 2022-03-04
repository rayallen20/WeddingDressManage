package sysError

import "strconv"

type LaundryRecordNotExistError struct {
	NotExistId int
}

func (l *LaundryRecordNotExistError) Error() string {
	if l.NotExistId != 0 {
		return "there doesn't exist laundryRecordInfo which id = " + strconv.Itoa(l.NotExistId)
	}
	return "laundryRecordInfo doesn't exist"
}
