package sysError

type DressHasUnavailableError struct {
	UnavailableStatus string
}

func (d *DressHasUnavailableError) Error() string {
	return "dress has " + d.UnavailableStatus + "ed"
}
