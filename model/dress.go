package model

import (
	"WeddingDressManage/lib/db"
	"errors"
	"gorm.io/gorm"
	"time"
)

// DressStatus 礼服状态
// onSale:在售
// preRent:预租
// rentOut:租借中
// pending:回收后待处理
// preOnSale:预上架
// laundry:送洗中
// maintain:维护中
// discard:已销库
// gift:已赠与
var DressStatus = map[string]string{
	"onSale":    "onSale",
	"preRent":   "preRent",
	"rentOut":   "rentOut",
	"pending":   "pending",
	"preOnSale": "preOnSale",
	"laundry":   "laundry",
	"maintain":  "maintain",
	"discard":   "discard",
	"gift":      "gift",
}

// Dress dress表的ORM
type Dress struct {
	Id              int
	CategoryId      int
	Category        *DressCategory `gorm:"foreignKey:CategoryId"`
	SerialNumber    int
	Size            string
	RentCounter     int
	LaundryCounter  int
	MaintainCounter int
	CoverImg        string
	SecondaryImg    string `gorm:"text"`
	Status          string
	CreatedTime     time.Time `gorm:"autoCreateTime"`
	UpdatedTime     time.Time `gorm:"autoUpdateTime"`
}

// FindMaxSerialNumberByCategoryId 根据品类Id 查找该品类下 礼服序号的最大值
// TODO:gorm没有max函数的实现吗?
func (d *Dress) FindMaxSerialNumberByCategoryId() (maxSerialNumber int, err error) {
	dresses := make([]*Dress, 0)
	err = db.Db.Where("category_id", d.CategoryId).Order("serial_number desc").Find(&dresses).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dresses[0].SerialNumber, nil
	}
	return 0, err
}

// AddDressesAndUpdateCategory 使用事务添加礼服同时更新礼服品类信息
func (d *Dress) AddDressesAndUpdateCategory(categoryORM *DressCategory, dressORMs []*Dress) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Updates(categoryORM).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(dressORMs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// FindUsableByCategoryId 查找指定品类下的可用礼服信息
// 可用:礼服状态不为已赠与 且 不为已销库的状态 即为可用状态
func (d *Dress) FindUsableByCategoryId(currentPage, itemPerPage int) (dresses []*Dress, err error) {
	dresses = make([]*Dress, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Not("status", []string{DressStatus["discard"], DressStatus["gift"]}).
		Where(d).Preload("Category").Preload("Category.Kind").Find(&dresses).Error
	return dresses, err
}

func (d *Dress) FindUnusableByCategoryId(currentPage, itemPerPage int) (dresses []*Dress, err error) {
	dresses = make([]*Dress, 0, itemPerPage)
	err = db.Db.Scopes(db.Paginate(currentPage, itemPerPage)).Where("status", []string{DressStatus["discard"], DressStatus["gift"]}).
		Where(d).Preload("Category").Preload("Category.Kind").Find(&dresses).Error
	return dresses, err
}

// CountUsableByCategoryId 统计指定品类下可用礼服的数量
func (d *Dress) CountUsableByCategoryId() (count int64, err error) {
	err = db.Db.Not("status", []string{DressStatus["discard"], DressStatus["gift"]}).Where(d).Find(d).Count(&count).Error
	return count, err
}

func (d *Dress) CountUnusableByCategoryId() (count int64, err error) {
	err = db.Db.Where("status", []string{DressStatus["discard"], DressStatus["gift"]}).Where(d).Find(d).Count(&count).Error
	return count, err
}

func (d *Dress) FindById() error {
	return db.Db.Where(d).Preload("Category").Preload("Category.Kind").First(d).Error
}

// UpdateDressStatusAndCreateLaundryRecord 使用事务修改礼服状态同时创建礼服送洗记录信息
func (d *Dress) UpdateDressStatusAndCreateLaundryRecord(categoryOrm *DressCategory, laundryRecordOrm *LaundryRecord) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Updates(d).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Updates(categoryOrm).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Omit("end_laundry_date").Create(laundryRecordOrm).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateDressStatusAndCreateMaintainRecord 使用事务修改礼服状态同时创建礼服维护记录信息
func (d *Dress) UpdateDressStatusAndCreateMaintainRecord(categoryOrm *DressCategory, maintainRecordOrm *MaintainRecord) error {
	tx := db.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Updates(d).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Updates(categoryOrm).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Select("source", "dress_id", "maintain_position_img", "note",
		"start_maintain_date", "plan_end_maintain_date", "status", "created_time",
		"updated_time").Create(maintainRecordOrm).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d *Dress) Updates() error {
	return db.Db.Updates(d).Error
}
