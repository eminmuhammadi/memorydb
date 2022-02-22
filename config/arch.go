package config

import (
	"gorm.io/gorm"
)

type MainArchitecture struct {
	DB *gorm.DB
}
