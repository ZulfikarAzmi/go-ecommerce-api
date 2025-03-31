package handlers

import (
    "fmt"
    "path/filepath"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "go-ecommerce-api/config"
    "go-ecommerce-api/database"
    "go-ecommerce-api/models"
    "time"
)

func UploadFotoProduk(c *fiber.Ctx) error {
    // Get user dari token JWT
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(*config.JWTClaim)
    userID := claims.ID

    // Ambil ID produk dari parameter
    productID := c.Params("id")

    // Cek apakah produk ada dan milik toko user yang login
    var product models.Product
    if err := database.DB.Preload("Toko").First(&product, productID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Produk tidak ditemukan",
        })
    }

    // Pastikan produk milik toko user yang login
    var toko models.Toko
    if err := database.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "message": "Anda tidak memiliki akses ke produk ini",
        })
    }

    if product.IDToko != toko.ID {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "message": "Produk ini bukan milik toko Anda",
        })
    }

    // Handle file upload
    file, err := c.FormFile("foto")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Gagal mengupload file",
            "error":   err.Error(),
        })
    }

    // Generate unique filename
    ext := filepath.Ext(file.Filename)
    filename := fmt.Sprintf("product_%d_%d%s", product.ID, time.Now().Unix(), ext)
    
    // Save file
    if err := c.SaveFile(file, fmt.Sprintf("./uploads/products/%s", filename)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal menyimpan file",
            "error":   err.Error(),
        })
    }

    // Simpan informasi foto ke database
    fotoProduk := models.FotoProduk{
        IDProduk:  product.ID,
        URL:       fmt.Sprintf("/uploads/products/%s", filename),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Simpan data foto ke database
    if err := database.DB.Create(&fotoProduk).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal menyimpan data foto",
            "error":   err.Error(),
        })
    }

    // Kembalikan respons dengan pesan sederhana
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Foto produk berhasil diupload",
    })
} 