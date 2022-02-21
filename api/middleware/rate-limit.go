package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func EnableRateLimit(app *fiber.App) fiber.Router {
	// Rate Limit
	return app.Use(limiter.New(limiter.Config{
		Max:               300,
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}
