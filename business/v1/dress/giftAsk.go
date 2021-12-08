package dress

import (
	"WeddingDressManage/business/v1/customer"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
)

type GiftAsk struct {
	Dress    *Dress
	Customer *customer.Customer
	Note     string
}

func (g *GiftAsk) Apply() error {
	orm := &model.DressGiftRecord{
		DressCategoryId: g.Dress.CategoryId,
		DressId:         g.Dress.Id,
		CustomerId:      g.Customer.Id,
		Note:            g.Note,
		Status:          model.GiftRecordStatus["pendingApproval"],
	}

	err := orm.Save()
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	return nil
}
