package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "go-ecommerce-api/config"
    "go-ecommerce-api/database"
    "go-ecommerce-api/models"
    "time"
)

// AddProduct menambahkan produk baru
func AddProduct(c *fiber.Ctx) error {
    // Get user dari token JWT
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(*config.JWTClaim)
    userID := claims.ID

    // Cek apakah user memiliki toko
    var toko models.Toko
    if err := database.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Anda harus memiliki toko untuk menambahkan produk",
        })
    }

    product := new(models.Product)
    if err := c.BodyParser(product); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Format request tidak valid",
            "error":   err.Error(),
        })
    }

    // Validasi input
    if product.NamaProduk == "" || product.HargaReseller == "" || 
       product.HargaKonsumen == "" || product.IDCategory == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Semua field wajib harus diisi",
        })
    }

    // Set ID toko dari user yang sedang login
    product.IDToko = toko.ID
    product.CreatedAt = time.Now()
    product.UpdatedAt = time.Now()

    // Cek apakah kategori exist
    var category models.Category
    if err := database.DB.First(&category, product.IDCategory).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Kategori tidak ditemukan",
        })
    }

    // Simpan produk ke database
    result := database.DB.Create(&product)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal menambahkan produk",
            "error":   result.Error.Error(),
        })
    }

    // Load relasi tanpa memuat user di dalam toko
    database.DB.Preload("Toko").Preload("Category").First(&product, product.ID)

    // Kembalikan respons dengan pesan sederhana
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Produk berhasil ditambahkan",
    })
}

// GetAllProducts mengambil semua produk
func GetAllProducts(c *fiber.Ctx) error {
    var products []models.Product
    result := database.DB.Preload("Toko").Preload("Category").Find(&products)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal mengambil data produk",
            "error":   result.Error.Error(),
        })
    }

    return c.JSON(products)
} 