package main

import (
	"fmt"

	db "github.com/eminmuhammadi/memorydb/db"
	model "github.com/eminmuhammadi/memorydb/model"
)

func main() {
	// Create new db connection
	conn, err := db.Init()

	// Check if database connection is valid
	if err != nil {
		panic(err)
	}

	// Create new user
	user := model.User{
		Name: "Emin Muhammadi",
	}

	// Save user to database
	if error := conn.Create(&user).Error; error != nil {
		panic(error)
	}

	// Get user from database
	if error := conn.First(&user, "name = ?", user.Name).Error; error != nil {
		panic(error)
	}

	// Print user id and creation time
	fmt.Println(user.Base.ID, user.Base.CreatedAt)

	// Close database connection
	defer db.Close(conn)
}
