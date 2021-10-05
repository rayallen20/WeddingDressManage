package dress

import (
	"WeddingDressManage/lib/structHelper"
	"WeddingDressManage/model"
	"time"
)

type Category struct {
	// 礼服品类ID
	Id int

	// 礼服类别ID
	KindId int

	// 礼服类别编码
	Code string

	// 礼服编号
	SerialNumber string

	// 可租礼服数量
	RentableQuantity int

	// 该品类礼服的总数量
	Quantity int

	// 该品类礼服总共被出租的次数
	RentNumber int

	// 该品类礼服总送洗次数
	LaundryNumber int

	// 租金
	RentMoney int

	// 押金
	CashPledge int

	// 状态
	Status string

	CreatedTime time.Time
	UpdatedTime time.Time
}

const (
	InitDressRentableQuantity = 1
	InitQuantity = 1
)

// Add 本方法用于在添加礼服时 确认品类信息
// 若待添加的礼服所属品类信息已存在 则将该品类的可租数量+1 总数量+1
// 若待添加礼服所属品类信息不存在 则创建品类信息 并设置租金/押金 置可租数量和总数量均为1
func (c *Category) Add(kindId, cashPledge, rentMoney int, code, serialNumber string) (err error) {
	var dressCategoryModel *model.DressCategory = &model.DressCategory{}

	// step1. 确认礼服品类信息是否存在
	err = dressCategoryModel.FindByKindIdAndCodeAndSN(kindId, code, serialNumber)
	if err != nil {
		return err
	}

	// step2. 品类信息不存在 创建品类信息
	if dressCategoryModel.Id == 0 {
		err = dressCategoryModel.Create(kindId, cashPledge, rentMoney, InitDressRentableQuantity, InitQuantity, code, serialNumber)
		// TODO:此处 c = (*Category)(dressCategoryModel) 赋值失败 为什么?
		structHelper.StructAssign(c, dressCategoryModel)
		return err
	}

	// step3. 品类信息存在 将该品类礼服的可租数量+1 总数量+1
	dressCategoryModel.RentableQuantity += 1
	dressCategoryModel.Quantity += 1
	err = dressCategoryModel.Save()
	structHelper.StructAssign(c, dressCategoryModel)
	return err
}


