package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var CustomerStatus map[string]string = map[string]string{
	"normal": "normal",
	"banned": "banned",
}

type Customer struct {
	Id           int
	Name         string
	Mobile       string
	Status       string
	BannedReason string    `gorm:"text"`
	CreatedTime  time.Time `gorm:"autoCreateTime"`
	UpdatedTime  time.Time `gorm:"autoUpdateTime"`
}

func (c *Customer) FindNormalByNameAndMobile() error {
	c.Status = CustomerStatus["normal"]
	return db.Db.Where(c).First(c).Error
}
