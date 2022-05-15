package customer

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"errors"
	"gorm.io/gorm"
)

type Customer struct {
	Id           int
	Name         string
	Mobile       string
	Status       string
	BannedReason string
}

func (c *Customer) FindNormalByNameAndMobile() error {
	orm := &model.Customer{
		Name:   c.Name,
		Mobile: c.Mobile,
	}
	err := orm.FindNormalByNameAndMobile()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.CustomerNotExistError{
			Name:   c.Name,
			Mobile: c.Mobile,
		}
	}

	c.fill(orm)
	return nil
}

func (c *Customer) fill(orm *model.Customer) {
	c.Id = orm.Id
	c.Name = orm.Name
	c.Mobile = orm.Mobile
	c.Status = orm.Status
	c.BannedReason = orm.BannedReason
}

func (c *Customer) FindOrCreateUser() error {
	err := c.FindNormalByNameAndMobile()
	var customerNotExistError *sysError.CustomerNotExistError
	if err != nil && !errors.As(err, &customerNotExistError) {
		return err
	}

	if errors.As(err, &customerNotExistError) {
		orm := &model.Customer{
			Name:   c.Name,
			Mobile: c.Mobile,
			Status: model.CustomerStatus["normal"],
		}
		err = orm.Save()
		if err != nil {
			return err
		}

		c.fill(orm)
	}

	return nil
}
