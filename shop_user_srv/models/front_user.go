package models

import "time"

type FrontUser struct {
	Id         int
	Email      string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (FrontUser) TableName() string {
	return "front_user"
}
