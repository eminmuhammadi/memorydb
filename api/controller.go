package api

import (
	model "github.com/eminmuhammadi/memorydb/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRoutes(server *fiber.App, sql *gorm.DB) {
	server.Get("/users", func(ctx *fiber.Ctx) error {
		return ctx.JSON(GetAllUsers(sql))
	})

	server.Post("/users", func(ctx *fiber.Ctx) error {
		user := new(model.User)

		if err := ctx.BodyParser(&user); err != nil {
			return ErrorHandler(ctx, err)
		}

		if user.Name == "" {
			return ctx.Status(400).SendString("Name is required")
		}

		return ctx.JSON(CreateUser(sql, user))
	})
}
