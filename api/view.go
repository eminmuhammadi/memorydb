package api

import (
	db "github.com/eminmuhammadi/memorydb/db"
	model "github.com/eminmuhammadi/memorydb/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserView struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var sql *gorm.DB = db.Init()

// Create User
func CreateUser(name string) UserView {
	// Create new user
	user := model.User{
		Name: name,
	}

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
func AllUsers() []UserView {
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
