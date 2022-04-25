package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// Img img表的ORM
type Img struct {
	Id              int
	SourceName      string
	DestinationName string
	Url             string
	CreatedTime     time.Time `gorm:"autoCreateTime"`
	UpdatedTime     time.Time `gorm:"autoUpdateTime"`
}

func (i *Img) Save() error {
	return db.Db.Save(i).Error
}
