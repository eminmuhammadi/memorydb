package main

import (
	"log"

	app "github.com/eminmuhammadi/memorydb/api"
	middleware "github.com/eminmuhammadi/memorydb/api/middleware"
	config "github.com/eminmuhammadi/memorydb/config"
	db "github.com/eminmuhammadi/memorydb/db"
	"gorm.io/gorm"
)

func main() {
	// In memory database
	var sql *gorm.DB = db.Init()

	// Migrate database
	if err := db.Migrate(sql); err != nil {
		panic(err)
	}

	// Create server
	server := app.CreateServer()

	// Middlewares
	middleware.EnableRequestID(server)
	middleware.EnableCors(server)
	middleware.EnableRecover(server)
	middleware.EnableEtag(server)
	middleware.EnableCompress(server)
	middleware.EnableRateLimit(server)
	middleware.EnableLogger(server)

	// Create routes
	app.CreateRoutes(server, sql)

	// Listen and serve on
	if error := app.StartServer(server, config.HOST, config.PORT); error != nil {
		log.Panic(error)
	}

	defer db.Close(sql)
}
