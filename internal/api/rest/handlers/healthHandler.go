package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvijaywargi/go-commerce/internal/api/rest"
)

type HealthHandler struct{}

func SetupHealthRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := HealthHandler{}

	app.Get("/health-check", handler.HealthCheck)
}

func (h *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"status":  "OK",
			"message": "Server is up and running",
		},
	)
}
