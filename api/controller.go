package api

import (
	"github.com/eminmuhammadi/memorydb/model"
	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(server *fiber.App) {
	server.Get("/users", func(ctx *fiber.Ctx) error {
		return ctx.JSON(AllUsers())
	})

	server.Post("/users", func(ctx *fiber.Ctx) error {
		user := new(model.User)

		if err := ctx.BodyParser(&user); err != nil {
			return ErrorHandler(ctx, err)
		}

		if user.Name == "" {
			return ctx.Status(400).SendString("Name is required")
		}

		return ctx.JSON(CreateUser(user.Name))
	})
}
