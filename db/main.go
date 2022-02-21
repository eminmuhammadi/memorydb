package db

import (
	"fmt"
	"time"

	config "github.com/eminmuhammadi/memorydb/config"
	date "github.com/eminmuhammadi/memorydb/date"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DSN string = fmt.Sprintf(":memory:?cache=%s&mode%s",
	"shared",
	"memory",
)

// Create database connection
func Create() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DSN), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NowFunc: func() time.Time {
			return date.Now(config.TimeZone)
		},
	})

	return db, err
}

// Intialize database in memory
func Init() *gorm.DB {
	db, err := Create()

	if err != nil {
		panic(err)
	}

	// Migrate database
	if err := Migrate(db); err != nil {
		panic(err)
	}

	return db
}

// Close database connection
func Close(db *gorm.DB) error {
	sqlDB, _ := db.DB()

	return sqlDB.Close()
}
