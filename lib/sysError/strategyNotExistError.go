package sysError

type StrategyNotExistError struct {
	Type string
}

func (s *StrategyNotExistError) Error() string {
	return s.Type + " strategy is not exist"
}
