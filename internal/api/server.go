package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamvijaywargi/go-commerce/configs"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
