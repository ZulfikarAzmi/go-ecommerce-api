package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	// Mengambil semua user dari database
	result := database.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data user",
			"error":   result.Error.Error(),
		})
	}

	// Mengembalikan data user dalam format JSON
	return c.JSON(users)
}
