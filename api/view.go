package api

import (
	model "github.com/eminmuhammadi/memorydb/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserView struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Create User
func CreateUser(sql *gorm.DB, user *model.User) UserView {
	// Save user to database
	if error := sql.Create(&user).Error; error != nil {
		panic(error)
	}

	// Get user from database
	if error := sql.Last(&user, "name = ?", user.Name).Error; error != nil {
		panic(error)
	}

	return UserView{
		ID:   user.ID,
		Name: user.Name,
	}
}

// List all users
func GetAllUsers(sql *gorm.DB) []UserView {
	// Get all users from database
	var users []model.User
	if error := sql.Find(&users).Error; error != nil {
		panic(error)
	}

	var usersView []UserView
	for _, user := range users {
		usersView = append(usersView, UserView{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return usersView
}
