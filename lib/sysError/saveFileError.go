package sysError

type SaveFileError struct {
	RealErr error
}

func (s *SaveFileError) Error() string {
	return s.RealErr.Error()
}
