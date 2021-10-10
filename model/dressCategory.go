package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var CategoryStatus = map[string]string{
	"usable": "usable",
	"unusable": "unusable",
}

type DressCategory struct {
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

	// 平均租金
	AvgRentMoney int

	// 主图
	MainImg string

	// 副图
	SecondaryImg string

	// 状态
	Status string

	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (d *DressCategory) FindByCodeAndSN(code, serialNumber string) (err error) {
	d.SerialNumber = serialNumber
	d.Code = code
	res := db.Db.Where(d).Find(d)
	return res.Error
}

func (d *DressCategory) Create(kindId, cashPledge, rentMoney, rentableQuantity, quantity int, code, categorySerialNumber, mainImg, secondaryImg string) (err error) {
	d.KindId = kindId
	d.CashPledge = cashPledge
	d.RentMoney = rentMoney
	d.Code = code
	d.SerialNumber = categorySerialNumber
	d.RentableQuantity = rentableQuantity
	d.Quantity = quantity
	d.MainImg = mainImg
	d.SecondaryImg = secondaryImg
	d.Status = CategoryStatus["usable"]
	res := db.Db.Create(d)
	return res.Error
}

func (d *DressCategory) Save() (err error) {
	res := db.Db.Save(d)
	return res.Error
}