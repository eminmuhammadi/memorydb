package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func EnableEtag(app *fiber.App) fiber.Router {
	// Etag
	return app.Use(etag.New(etag.Config{
		Weak: true,
	}))
}
