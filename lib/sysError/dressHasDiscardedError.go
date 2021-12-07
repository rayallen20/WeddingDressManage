package sysError

type DressHasDiscardedError struct {
}

func (d *DressHasDiscardedError) Error() string {
	return "dress has " + "discarded"
}
