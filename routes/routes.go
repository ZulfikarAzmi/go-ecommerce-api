package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/handlers"
)

func Setup(app *fiber.App) {
	// ... existing routes ...
	
	// Auth routes
	app.Post("/api/login", handlers.Login)
	app.Post("/api/register", handlers.AddUser)
} 