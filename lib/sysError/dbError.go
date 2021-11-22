package sysError

// DbError 数据库错误
type DbError struct {
	RealError error
}

func (d *DbError) Error() string {
	return "DbError: " + d.RealError.Error()
}
