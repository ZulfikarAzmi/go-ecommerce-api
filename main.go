package main

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/handlers"
	"log"
	"os"
)

func main() {
	// Koneksi ke database
	database.Connect()

	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Endpoints
	app.Get("/users", handlers.GetAllUsers)
	app.Get("/users/:id", handlers.GetUserByID)
	app.Post("/users", handlers.AddUser)
	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("E-Commerce API is running!")
	})

	log.Fatal(app.Listen(":" + port))
}
