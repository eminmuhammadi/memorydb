package db

import (
	"time"

	config "github.com/eminmuhammadi/memorydb/config"
	date "github.com/eminmuhammadi/memorydb/date"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Create database connection
func Create() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:?cache=shared&mode=memory"), &gorm.Config{
		NowFunc: func() time.Time {
			return date.Now(config.TimeZone)
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

// Intialize database in memory
func Init() (*gorm.DB, error) {
	db, err := Create()

	if err != nil {
		return nil, err
	}

	// Migrate database
	if err := Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// Close database connection
func Close(db *gorm.DB) error {
	sqlDB, _ := db.DB()

	return sqlDB.Close()
}
