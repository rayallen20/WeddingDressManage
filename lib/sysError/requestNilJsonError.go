package sysError

type RequestNilJsonError struct {
}

func (r *RequestNilJsonError) Error() string {
	return "request param is nil json"
}
