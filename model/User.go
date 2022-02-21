package model

import (
	config "github.com/eminmuhammadi/memorydb/config"
	date "github.com/eminmuhammadi/memorydb/date"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Name string    `gorm:"type:varchar(128);not null;" json:"name"`
}

// Before first time creation
func (user *User) BeforeCreate(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Set UUID for ID
	user.ID = uuid.NewV4()

	// Set timestamp for base
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	return nil
}

// Before update
func (user *User) BeforeUpdate(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Update timestamp for updated_at
	user.UpdatedAt = currentTime

	return nil
}

// Before delete
func (user *User) BeforeDelete(tx *gorm.DB) error {
	currentTime := date.Now(config.TimeZone)

	// Set timestamp for deleted_at
	user.DeletedAt.Time = currentTime

	return nil
}
