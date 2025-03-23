package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers retrieves all users from the database.
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println("Error retrieving users:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve users",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(users)
}

// GetUserByID retrieves a user by ID from the database.
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(user)
}

// AddUser menambahkan user baru (register) ke database
func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Validasi input
	if user.Nama == "" || user.KataSandi == "" || user.Notelp == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama, kata sandi, nomor telepon, dan email harus diisi",
		})
	}

	// Validasi email unik
	var existingUser models.User
	if result := database.DB.Where("email = ?", user.Email).First(&existingUser); result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email sudah terdaftar",
		})
	}

	// Validasi nomor telepon unik
	if result := database.DB.Where("notelp = ?", user.Notelp).First(&existingUser); result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nomor telepon sudah terdaftar",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memproses kata sandi",
		})
	}
	user.KataSandi = string(hashedPassword)

	// Set default values
	user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Simpan user ke database
	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendaftarkan user",
			"error":   result.Error.Error(),
		})
	}

	// Hapus kata sandi dari response
	user.KataSandi = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"user":    user,
	})
}
