package sysError

type DateBeforeTodayError struct {
	Field string
}

func (d *DateBeforeTodayError) Error() string {
	return d.Field + " must be later than today"
}
