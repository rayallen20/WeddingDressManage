package sysError

type DressHasGiftedError struct {

}

func (d *DressHasGiftedError) Error() string {
	return "dress has " + "gifted"
}
