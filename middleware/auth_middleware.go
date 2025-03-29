package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-ecommerce-api/config"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get cookie
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Parse token
	token, err := jwt.ParseWithClaims(cookie, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Set user ke context
	c.Locals("user", token)
	
	return c.Next()
} 