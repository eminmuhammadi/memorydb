package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var APP_CONFIGURATION = fiber.Config{
	Prefork:                 true,
	StrictRouting:           true,
	CaseSensitive:           true,
	BodyLimit:               4 * 1024 * 1024,
	Concurrency:             256 * 1024,
	ReadTimeout:             time.Second * 15,
	WriteTimeout:            time.Second * 15,
	IdleTimeout:             time.Second * 60,
	ReadBufferSize:          4096,
	WriteBufferSize:         4096,
	ErrorHandler:            ErrorHandler,
	DisableStartupMessage:   true,
	AppName:                 "Memorydb v0.0.1",
	EnableTrustedProxyCheck: true,
	TrustedProxies: []string{
		"127.0.0.1",
	},
}

// CreateServer creates a new server instance
func CreateServer() *fiber.App {
	return fiber.New(APP_CONFIGURATION)
}

// Creates a listener
func StartServer(app *fiber.App, host string, port string) error {
	return app.Listen(fmt.Sprintf("%s:%s", host, port))
}
