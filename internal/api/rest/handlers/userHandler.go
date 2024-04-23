package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvijaywargi/go-commerce/internal/api/rest"
)

type UserHandler struct {
	//
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// Create an instance of user service and pass to user handler

	handler := UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)

	app.Get("/carts", handler.GetCart)
	app.Post("/carts", handler.AddToCart)

	app.Get("/orders", handler.GetOrders)
	app.Post("/orders", handler.CreateOrder)
	app.Get("/orders/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Register",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GetVerificationCode",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GetProfile",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Verify",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "CreateProfile",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GetCart",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "AddToCart",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "CreateOrder",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GetOrders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GetOrder",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "BecomeSeller",
	})
}
