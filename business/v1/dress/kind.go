package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"errors"
	"gorm.io/gorm"
)

// Kind 礼服大类 如同超市商品分类中的 "食品" "日用品" 等类别
type Kind struct {
	Id int
	Name string
	Code string
	Status string
}

// FindById 根据Id属性值查找礼服大类信息
func (k *Kind) FindById() error {
	kindModel := &model.DressKind{
		Id: k.Id,
	}
	err := kindModel.FindById()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			kindNotExistError := &sysError.KindNotExistError{NotExistId: k.Id}
			return kindNotExistError
		}
		return err
	}

	k.Code = kindModel.Code
	k.Name = kindModel.Name
	k.Status = kindModel.Status
	return nil
}