package main

import (
	"fmt"

	config "github.com/eminmuhammadi/memorydb/config"
	db "github.com/eminmuhammadi/memorydb/db"
	model "github.com/eminmuhammadi/memorydb/model"
	view "github.com/eminmuhammadi/memorydb/view"
)

func main() {
	// App structure
	server := &config.MainArchitecture{
		DB: db.Init(),
	}

	// Migrate database
	if err := db.Migrate(server.DB); err != nil {
		panic(err)
	}

	view.CreateUser(server.DB, &model.User{
		Name: "John Doe",
	})

	users := view.GetAllUsers(server.DB)
	fmt.Println(users)

	defer db.Close(server.DB)
}
