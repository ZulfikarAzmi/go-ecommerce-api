package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-ecommerce-api/database"
	"go-ecommerce-api/handlers"
	"go-ecommerce-api/middleware"
	"log"
)

func main() {
	// Inisialisasi koneksi database
	database.Connect()

	app := fiber.New()

	// Middleware CORS dengan konfigurasi yang aman
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Sesuaikan dengan origin frontend Anda
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Grup untuk API
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.AddUser)
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)

	// Protected routes
	api.Get("/welcome", middleware.AuthMiddleware, handlers.Welcome)

	// User routes
	users := api.Group("/users")
	users.Get("/", handlers.GetAllUsers)
	users.Get("/:id", handlers.GetUserByID)

	// Toko routes
	toko := api.Group("/toko")
	toko.Get("/", handlers.GetAllToko)        // Route untuk mendapatkan semua toko
	toko.Get("/:id", handlers.GetTokoByID)    // Route untuk mendapatkan toko berdasarkan ID

	// Alamat routes (protected, perlu login)
	alamat := api.Group("/alamat")
	alamat.Use(middleware.AuthMiddleware) // Middleware untuk memastikan user sudah login
	alamat.Post("/", handlers.AddAlamat)
	alamat.Get("/", handlers.GetUserAlamat)
	alamat.Get("/:id", handlers.GetAlamatByID)
	alamat.Put("/:id", handlers.UpdateAlamat)
	alamat.Delete("/:id", handlers.DeleteAlamat)

	// Category routes
	category := api.Group("/categories")
	category.Get("/", handlers.GetAllCategories)                                    // Dapat diakses semua user
	category.Post("/", middleware.AuthMiddleware, middleware.AdminMiddleware, handlers.AddCategory)  // Hanya admin

	// Product routes
	product := api.Group("/products")
	product.Get("/", handlers.GetAllProducts)                           // Dapat diakses semua user
	product.Post("/", middleware.AuthMiddleware, handlers.AddProduct)   // Hanya pemilik toko

	// Start server
	log.Fatal(app.Listen(":8080"))
}
