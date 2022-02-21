package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func EnableRecover(app *fiber.App) fiber.Router {
	// Recover
	return app.Use(recover.New())
}
