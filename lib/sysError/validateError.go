package sysError

// ValidateError 参数校验失败错误
type ValidateError struct {
	// Key 校验失败的字段
	Key string
	// Msg 校验失败信息
	Msg string
}

func (v ValidateError) Error() string {
	return v.Key + ":" + v.Msg
}