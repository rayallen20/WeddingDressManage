package category

import (
	"WeddingDressManage/business/v1/dress/unit"
	"WeddingDressManage/lib/sliceHelper"
	"WeddingDressManage/model"
)

type Category struct {
	// 礼服品类ID
	Id int `json:"id"`

	// 礼服类别ID
	KindId int `json:"kindId"`

	// 礼服类别编码
	Code string `json:"code"`

	// 礼服品类编号
	SerialNumber string `json:"serialNumber"`

	// 可租礼服数量
	RentableQuantity int `json:"rentableQuantity"`

	// 该品类礼服的总数量
	Quantity int `json:"quantity"`

	// 租金
	CharterMoney int `json:"charterMoney"`

	// 押金
	CashPledge int `json:"cashPledge"`

	// 该品类礼服总共被出租的次数
	RentNumber int `json:"rentNumber"`

	// 该品类礼服总送洗次数
	LaundryNumber int `json:"laundryNumber"`

	// 平均租金
	AvgRentMoney int `json:"avgRentMoney"`

	// 封面图
	CoverImg string `json:"coverImg"`

	// 副图
	SecondaryImg []string `json:"secondaryImg,omitempty"`

	// 状态
	Status string `json:"status,omitempty"`
}

// Add 添加品类的同时 添加多件礼服
// 业务逻辑上不允许在添加一个品类时 不添加具体的礼服
func (c *Category) Add(units []*unit.Unit) ([]*unit.Unit, error) {
	categoryModel := &model.DressCategory{
		KindId: c.KindId,
		Code: c.Code,
		SerialNumber: c.SerialNumber,
		RentableQuantity: c.RentableQuantity,
		Quantity: c.Quantity,
		CharterMoney: c.CharterMoney,
		CashPledge: c.CashPledge,
		CoverImg: c.CoverImg,
		SecondaryImg: sliceHelper.ConvertStrSliceToStr(c.SecondaryImg, "|"),
		Status: model.CategoryStatus["usable"],
	}

	unitModels := make([]*model.DressUnit, 0, len(units))
	for i := 0; i < len(units); i++ {
		unitModel := &model.DressUnit{
			SerialNumber: units[i].SerialNumber,
			Size: units[i].Size,
			CoverImg: units[i].CoverImg,
			SecondaryImg: sliceHelper.ConvertStrSliceToStr(units[i].SecondaryImg, "|"),
			Status:model.UnitStatus["rentable"],
		}
		unitModels = append(unitModels, unitModel)
	}

	unitModels, err := categoryModel.AddCategoryAndUnits(unitModels)
	if err != nil {
		return nil, err
	}

	c.Id = categoryModel.Id
	for i := 0; i < len(unitModels); i++ {
		units[i].Id = unitModels[i].Id
		units[i].CategoryId = c.Id
	}
	return units, nil
}

// FindByKindIdAndCodeAndSN 根据类别编码和品类编号查询信息
func (c *Category) FindByKindIdAndCodeAndSN() (err error) {
	model := &model.DressCategory{}
	err = model.FindByKindIdAndCodeAndSN(c.KindId, c.Code, c.SerialNumber)
	if err != nil {
		return err
	}
	c.Id = model.Id
	return nil
}