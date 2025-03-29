package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-ecommerce-api/config"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"time"
)

// AddAlamat menambahkan alamat baru untuk user yang sedang login
func AddAlamat(c *fiber.Ctx) error {
	// Get user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	alamat := new(models.Alamat)
	if err := c.BodyParser(alamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Set ID user dari token
	alamat.IDUser = userID
	alamat.CreatedAt = time.Now()
	alamat.UpdatedAt = time.Now()

	// Validasi input
	if alamat.JudulAlamat == "" || alamat.NamaPenerima == "" || alamat.NoTelp == "" || alamat.DetailAlamat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Semua field harus diisi",
		})
	}

	// Ambil data user terlebih dahulu
	var userData models.User
	if err := database.DB.First(&userData, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data user",
			"error":   err.Error(),
		})
	}

	// Set data user ke alamat
	alamat.User = userData

	// Simpan alamat ke database
	result := database.DB.Create(&alamat)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan alamat",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Alamat berhasil ditambahkan",
		"alamat":  alamat,
	})
}

// GetUserAlamat mengambil semua alamat milik user yang sedang login
func GetUserAlamat(c *fiber.Ctx) error {
	// Get user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	var alamatList []models.Alamat
	// Gunakan Preload untuk mengambil data user
	result := database.DB.Preload("User").Where("id_user = ?", userID).Find(&alamatList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data alamat",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(alamatList)
}

// GetAlamatByID mengambil alamat berdasarkan ID (hanya untuk alamat milik user yang login)
func GetAlamatByID(c *fiber.Ctx) error {
	// Get user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	id := c.Params("id")
	var alamat models.Alamat

	// Gunakan Preload untuk mengambil data user
	result := database.DB.Preload("User").Where("id = ? AND id_user = ?", id, userID).First(&alamat)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Alamat tidak ditemukan",
		})
	}

	return c.JSON(alamat)
}

// UpdateAlamat mengupdate alamat berdasarkan ID
func UpdateAlamat(c *fiber.Ctx) error {
	// Get user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	id := c.Params("id")
	var alamat models.Alamat

	// Cari alamat dan pastikan milik user yang sedang login
	result := database.DB.Preload("User").Where("id = ? AND id_user = ?", id, userID).First(&alamat)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Alamat tidak ditemukan",
		})
	}

	// Parse body request
	updateData := new(models.Alamat)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Update data alamat
	alamat.JudulAlamat = updateData.JudulAlamat
	alamat.NamaPenerima = updateData.NamaPenerima
	alamat.NoTelp = updateData.NoTelp
	alamat.DetailAlamat = updateData.DetailAlamat
	alamat.UpdatedAt = time.Now()

	// Simpan perubahan
	database.DB.Save(&alamat)

	return c.JSON(fiber.Map{
		"message": "Alamat berhasil diupdate",
		"alamat": alamat,
	})
}

// DeleteAlamat menghapus alamat berdasarkan ID (hanya untuk alamat milik user yang login)
func DeleteAlamat(c *fiber.Ctx) error {
	// Get user dari token JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)
	userID := claims.ID

	id := c.Params("id")
	var alamat models.Alamat

	// Cari alamat dan pastikan milik user yang sedang login
	result := database.DB.Where("id = ? AND id_user = ?", id, userID).First(&alamat)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Alamat tidak ditemukan",
		})
	}

	// Hapus alamat
	database.DB.Delete(&alamat)

	return c.JSON(fiber.Map{
		"message": "Alamat berhasil dihapus",
	})
} 