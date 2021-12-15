package sysError

type LaundryStatusError struct {
	DressNowStatus string
}

func (l *LaundryStatusError) Error() string {
	return "Can't laundry dress because the dress is " + l.DressNowStatus + " now."
}
