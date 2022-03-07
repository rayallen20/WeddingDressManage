package sysError

import "strconv"

type DressIsNotMaintainingError struct {
	NotMaintainingId int
}

func (d *DressIsNotMaintainingError) Error() string {
	if d.NotMaintainingId != 0 {
		return "there doesn't exist maintaining dress which id = " + strconv.Itoa(d.NotMaintainingId)
	}
	return "maintaining dress doesn't exist"
}
