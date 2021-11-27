package sysError

import "strconv"

type CategoryNotExistError struct {
	Id int
}

func (c *CategoryNotExistError) Error() string {
	if c.Id != 0 {
		return "the categoryInfo which id = " + strconv.Itoa(c.Id) + " doesn't exist"
	}
	return "the categoryInfo doesn't exist"
}
