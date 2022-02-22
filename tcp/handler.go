package tcp

import (
	"fmt"

	"github.com/eminmuhammadi/memorydb/model"
	"github.com/eminmuhammadi/memorydb/view"
	"gorm.io/gorm"
)

// Query
func Handler(sql *gorm.DB, query string) string {
	// Create new user using query
	view.CreateUser(sql, &model.User{
		Name: query,
	})

	// Get all users
	users := view.GetAllUsers(sql)

	// Return all users
	return fmt.Sprintln(users)
}
