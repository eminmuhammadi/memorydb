package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Base

	Name string `gorm:"type:varchar(128);not null;"`
}
