package main

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/handlers"
	"log"
	"os"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Koneksi ke database
	database.Connect()

	// Mendapatkan port dari environment atau menggunakan default 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Route untuk mendapatkan semua user
	app.Get("/users", handlers.GetUsers)

	// Route default
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("E-Commerce API is running!")
	})

	// Menjalankan server
	log.Fatal(app.Listen(":" + port))
}
