package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-ecommerce-api/config"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
)

func AdminMiddleware(c *fiber.Ctx) error {
	// Ambil user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	// Cek apakah user adalah admin
	var userData models.User
	if err := database.DB.First(&userData, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data user",
		})
	}

	if !userData.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak. Hanya admin yang dapat mengakses.",
		})
	}

	return c.Next()
} 