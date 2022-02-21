package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func EnableCompress(app *fiber.App) fiber.Router {
	// Compress
	return app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
}
