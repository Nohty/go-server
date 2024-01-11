package router

import (
	"github.com/Nohty/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Status)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
}
