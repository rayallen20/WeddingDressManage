package unit

import (
	"WeddingDressManage/model"
	"strings"
)

type Unit struct {
	// 礼服ID
	Id int	`json:"id"`

	// 礼服品类ID
	CategoryId int `json:"categoryId"`

	// 礼服序号
	SerialNumber int `json:"serialNumber"`

	// 尺码
	Size string `json:"size"`

	// 出租次数
	RentNumber int `json:"rentNumber"`

	// 送洗次数
	LaundryNumber int `json:"laundryNumber"`

	// 封面图
	CoverImg string `json:"coverImg"`

	// 副图
	SecondaryImg []string `json:"secondaryImg"`

	// 状态
	Status string `json:"status"`
}

func (u Unit) ShowUsable(categoryId int, page int) (units []Unit, err error) {
	unitModel := model.DressUnit{}
	status := []string{
		model.UnitStatus["rentable"],
		model.UnitStatus["rentOut"],
		model.UnitStatus["laundry"],
	}

	unitInfos, err := unitModel.FindByCategoryIdAndStatus(categoryId, page, status)
	if err != nil {
		return nil, err
	}

	units = make([]Unit, 0, len(unitInfos))
	for i := 0; i < len(unitInfos); i++ {
		unit := Unit{
			Id:            unitInfos[i].Id,
			CategoryId:    categoryId,
			SerialNumber:  unitInfos[i].SerialNumber,
			Size:          unitInfos[i].Size,
			RentNumber:    unitInfos[i].RentNumber,
			LaundryNumber: unitInfos[i].LaundryNumber,
			CoverImg:      unitInfos[i].CoverImg,
			SecondaryImg:  strings.Split(unitInfos[i].SecondaryImg, "|"),
			Status:        unitInfos[i].Status,
		}

		units = append(units, unit)
	}

	return units, nil
}

func (u Unit) CountCategoryUsable(categoryId int) (int64, error) {
	unitModel := model.DressUnit{}
	return unitModel.CountUsableByCategoryId(categoryId)
}

func (u Unit) ShowUnusable(categoryId int, page int) (units []Unit, err error) {
	unitModel := model.DressUnit{}
	status := []string{
		model.UnitStatus["obsolete"],
		model.UnitStatus["gift"],
	}

	unitInfos, err := unitModel.FindByCategoryIdAndStatus(categoryId, page, status)
	if err != nil {
		return nil, err
	}

	units = make([]Unit, 0, len(unitInfos))
	for i := 0; i < len(unitInfos); i++ {
		unit := Unit{
			Id:            unitInfos[i].Id,
			CategoryId:    categoryId,
			SerialNumber:  unitInfos[i].SerialNumber,
			Size:          unitInfos[i].Size,
			RentNumber:    unitInfos[i].RentNumber,
			LaundryNumber: unitInfos[i].LaundryNumber,
			CoverImg:      unitInfos[i].CoverImg,
			SecondaryImg:  strings.Split(unitInfos[i].SecondaryImg, "|"),
			Status:        unitInfos[i].Status,
		}

		units = append(units, unit)
	}

	return units, nil
}

func (u Unit) CountCategoryUnusable(categoryId int) (int64, error) {
	unitModel := model.DressUnit{}
	return unitModel.CountUnusableByCategoryId(categoryId)
}