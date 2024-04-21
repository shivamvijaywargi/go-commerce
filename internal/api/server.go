package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvijaywargi/go-commerce/configs"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	app.Get("/health-check", HealthCheck)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is up and running",
	})
}
