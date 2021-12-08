package sysError

type CustomerNotExistError struct {
	Name   string
	Mobile string
}

func (c *CustomerNotExistError) Error() string {
	if c.Name != "" && c.Mobile != "" {
		return "the customer which name = " + c.Name + " and mobile = " + c.Mobile + " doesn't exist"
	}

	return "customer doesn't exist"
}
