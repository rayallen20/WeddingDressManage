package sysError

type CustomerBeBannedError struct {
	Name   string
	Mobile string
}

func (c *CustomerBeBannedError) Error() string {
	return "The customer which name = " + c.Name + " and mobile = " + c.Mobile + " is be banned"
}
