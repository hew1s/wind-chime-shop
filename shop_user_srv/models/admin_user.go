package models

import "time"

type AdminUser struct {
	Id         int
	UserName   string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (AdminUser) AdminUser() string {
	return "admin_user"
}
