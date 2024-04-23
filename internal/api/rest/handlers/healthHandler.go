package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoute(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"status":  "OK",
			"message": "Server is up and running",
		},
	)
}
