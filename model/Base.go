package model

import (
	"time"

	config "github.com/eminmuhammadi/memorydb/config"
	date "github.com/eminmuhammadi/memorydb/date"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ref:https://gorm.io/docs/hooks.html
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp: not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Before first time creation
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Set UUID for ID
	base.ID = uuid.NewV4()

	// Set timestamp for base
	base.CreatedAt = currentTime
	base.UpdatedAt = currentTime

	return nil
}

// Before update
func (base *Base) BeforeUpdate(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Update timestamp for updated_at
	base.UpdatedAt = currentTime

	return nil
}

// Before delete
func (base *Base) BeforeDelete(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Set timestamp for deleted_at
	base.DeletedAt.Time = currentTime

	return nil
}
