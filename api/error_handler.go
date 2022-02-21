package api

import "github.com/gofiber/fiber/v2"

// Error handler
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(500).SendString(err.Error())
}
