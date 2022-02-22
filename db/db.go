package db

import (
	"fmt"
	"time"

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
func Create(timezone string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DSN), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Warn),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		QueryFields:            true,
		NowFunc: func() time.Time {
			return date.Now(timezone)
		},
	})

	return db, err
}

// Intialize database in memory
func Init(timezone string) *gorm.DB {
	db, err := Create(timezone)

	if err != nil {
		panic(err)
	}

	return db
}

// Close database connection
func Close(db *gorm.DB) error {
	sqlDB, _ := db.DB()

	return sqlDB.Close()
}
