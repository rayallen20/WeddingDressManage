package kind

import (
	"WeddingDressManage/lib/wdmError"
	"WeddingDressManage/model"
)

type Kind struct {
	// 主键自增ID
	Id int	`json:"id"`

	// 品类名称
	Name string `json:"name"`

	// 品类编码
	Code string `json:"code"`

	// 品类状态(可用/不可用)
	Status string `json:"status,omitempty"`
}

func (k *Kind) Show() (kinds []Kind, err error) {
	kindModel := &model.DressKind{}
	dressKinds, dbErr := kindModel.FindAllUsableKinds()
	if dbErr != nil {
		err = wdmError.DBError {
			Message: dbErr.Error(),
		}
		return nil, err
	}

	kinds = make([]Kind, 0, len(dressKinds))
	for _, dressKind := range dressKinds {
		kind := Kind{
			Id: dressKind.Id,
			Name: dressKind.Name,
			Code: dressKind.Code,
		}
		kinds = append(kinds, kind)
	}
	return kinds, nil
}

func (k *Kind) FindByNameAndCode() (err error) {
	kindModel := &model.DressKind{}
	err = kindModel.FindByKindAndCode(k.Name, k.Code)
	if err != nil {
		return err
	}
	k.Id = kindModel.Id
	return nil
}
