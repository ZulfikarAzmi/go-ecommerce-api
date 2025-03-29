package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"go-ecommerce-api/config"
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

	// Set default values untuk user
	user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Mulai transaksi database
	tx := database.DB.Begin()

	// Simpan user ke database
	if result := tx.Create(&user); result.Error != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendaftarkan user",
			"error":   result.Error.Error(),
		})
	}

	// Buat toko otomatis untuk user
	toko := &models.Toko{
		IDUser:    user.ID,
		NamaToko:  "Toko " + user.Nama, // Default nama toko
		URLFoto:   "",                   // Default kosong
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Simpan toko ke database
	if result := tx.Create(&toko); result.Error != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat toko",
			"error":   result.Error.Error(),
		})
	}

	// Commit transaksi
	tx.Commit()

	// Hapus kata sandi dari response
	user.KataSandi = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"user":    user,
		"toko":    toko,
	})
}

// Login handles user authentication and returns JWT token
func Login(c *fiber.Ctx) error {
	loginRequest := struct {
		Email     string `json:"email"`
		KataSandi string `json:"kata_sandi"`
	}{}

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Cari user berdasarkan email
	var user models.User
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email atau kata sandi salah",
		})
	}

	// Verifikasi kata sandi
	err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(loginRequest.KataSandi))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email atau kata sandi salah",
		})
	}

	// Buat token JWT
	claims := &config.JWTClaim{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat token",
			"error":   err.Error(),
		})
	}

	// Set cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"nama":  user.Nama,
		},
	})
}

// Logout handles user logout
func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour) // Set expired ke masa lalu
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}

// Welcome handles protected welcome route
func Welcome(c *fiber.Ctx) error {
	// Get user dari middleware
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaim)

	var userData models.User
	result := database.DB.First(&userData, claims.ID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Welcome, " + userData.Nama,
	})
}
