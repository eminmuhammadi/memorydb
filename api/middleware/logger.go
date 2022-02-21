package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func EnableLogger(app *fiber.App) fiber.Router {
	// Logger
	return app.Use(logger.New(logger.Config{
		Format:       "[${time}]\t${locals:requestid}\t${ip}\t${status}\t${method}\t${path}\t${latency}\n",
		TimeFormat:   "2006-01-02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stderr,
	}))
}
