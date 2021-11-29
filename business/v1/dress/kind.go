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
	orm := &model.DressKind{
		Id: k.Id,
	}
	err := orm.FindById()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			kindNotExistError := &sysError.KindNotExistError{NotExistId: k.Id}
			return kindNotExistError
		}
		return err
	}

	k.fill(orm)

	return nil
}

// fill 根据DressKind orm填充一个biz层的Kind对象
func (k *Kind) fill(orm *model.DressKind)  {
	k.Id = orm.Id
	k.Code = orm.Code
	k.Name = orm.Name
	k.Status = orm.Status
}

func (k *Kind) Show() ([]*Kind, error) {
	orm := &model.DressKind{}
	orms, err := orm.FindAllOnSale()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		dbErr := &sysError.DbError{RealError: err}
		return nil, dbErr
	}

	kinds := make([]*Kind, 0, len(orms))
	for _, kindOrm := range orms {
		kind := &Kind{}
		kind.fill(kindOrm)
		kinds = append(kinds, kind)
	}

	return kinds, nil
}