package sysError

type CategoryHasExistError struct {
	SerialNumber string
}

func (c *CategoryHasExistError) Error() string {
	if c.SerialNumber != "" {
		return "the categoryInfo which serialNumber = " + c.SerialNumber + " has exist"
	}
	return "categoryInfo has exist"
}
