package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"log"
)

// GetAllToko retrieves all toko from the database.
func GetAllToko(c *fiber.Ctx) error {
	var tokos []models.Toko

	result := database.DB.Find(&tokos)
	if result.Error != nil {
		log.Println("Error retrieving tokos:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve tokos",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(tokos)
}

// GetTokoByID retrieves a toko by ID from the database.
func GetTokoByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var toko models.Toko
	result := database.DB.First(&toko, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Toko not found",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(toko)
} 