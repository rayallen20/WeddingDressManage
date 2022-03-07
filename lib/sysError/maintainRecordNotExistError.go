package sysError

import "strconv"

type MaintainRecordNotExistError struct {
	NotExistId int
}

func (m *MaintainRecordNotExistError) Error() string {
	if m.NotExistId != 0 {
		return "there doesn't exist maintainRecordInfo which id = " + strconv.Itoa(m.NotExistId)
	}
	return "laundryRecordInfo doesn't exist"
}
