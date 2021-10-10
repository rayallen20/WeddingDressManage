package wdmError

type DBError struct {
	Code int
	Message string
}

func (e DBError) Error() string {
	return e.Message
}
