package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamvijaywargi/go-commerce/configs"
	"github.com/shivamvijaywargi/go-commerce/internal/api/rest"
	"github.com/shivamvijaywargi/go-commerce/internal/api/rest/handlers"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// Health check handler
	handlers.SetupHealthRoutes(rh)

	// user handler
	handlers.SetupUserRoutes(rh)

	// transaction

	// catalog
}
