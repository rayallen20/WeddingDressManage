package sysError

type MaintainStatusError struct {
	DressNowStatus string
}

func (m *MaintainStatusError) Error() string {
	return "Can't maintain dress because the dress is " + m.DressNowStatus + " now."
}
