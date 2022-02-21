package main

import (
	"log"

	app "github.com/eminmuhammadi/memorydb/api"
	middleware "github.com/eminmuhammadi/memorydb/api/middleware"
	config "github.com/eminmuhammadi/memorydb/config"
)

func main() {
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
	app.CreateRoutes(server)

	// Listen and serve on
	error := app.StartServer(server, config.HOST, config.PORT)

	if error != nil {
		log.Panic(error)
	}
}
