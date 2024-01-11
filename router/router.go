package router

import (
	"github.com/Nohty/api/handler"
	"github.com/Nohty/api/middleware"

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

	// User
	user := api.Group("/user")
	user.Get("/:id<int>", middleware.Protected(), handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id<int>", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id<int>", middleware.Protected(), handler.DeleteUser)
	user.Put("/:id<int>/password", middleware.Protected(), handler.UpdatePassword)
}
