package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"time"
)

// AddCategory menambahkan kategori baru (hanya admin)
func AddCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Validasi input
	if category.NamaCategory == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama kategori harus diisi",
		})
	}

	// Set waktu
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	// Simpan kategori ke database
	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan kategori",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Kategori berhasil ditambahkan",
		"category": category,
	})
}

// GetAllCategories mengambil semua kategori
func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data kategori",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(categories)
} 