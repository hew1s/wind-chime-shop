package models

import "time"

type SecKills struct {
	Id         int
	Name       string
	Price      float32 `gorm:"type:decimal(11,2)"`
	Num        int
	PId        int
	StartTime  time.Time
	EndTime    time.Time
	Status     int
	CreateTime time.Time
}

func (SecKills) TableName() string {
	return "product_seckills"
}
