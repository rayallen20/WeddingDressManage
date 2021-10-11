package dress

import (
	"WeddingDressManage/lib/wdmError"
	"WeddingDressManage/model"
)

type Kind struct {
	// 主键自增ID
	Id int	`json:"id"`

	// 品类名称
	Kind string `json:"kind"`

	// 品类编码
	Code string `json:"code"`

	// 品类状态(可用/不可用)
	Status string `json:"status,omitempty"`
}

func (d *Kind) Show() (kinds []Kind, err error) {
	dressKinds, dbErr := model.FindAllUsableKinds()
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
			Kind: dressKind.Kind,
			Code: dressKind.Code,
		}
		kinds = append(kinds, kind)
	}
	return kinds, nil
}
