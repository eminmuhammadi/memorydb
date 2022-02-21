package db

import (
	model "github.com/eminmuhammadi/memorydb/model"
	"gorm.io/gorm"
)

// Migrating data from source
func Migrate(conn *gorm.DB) error {
	error := conn.AutoMigrate(
		&model.User{},
	)

	if error != nil {
		return error
	}

	return nil
}
