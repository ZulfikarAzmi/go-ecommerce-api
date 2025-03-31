package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "go-ecommerce-api/config"
    "go-ecommerce-api/database"
    "go-ecommerce-api/models"
    "time"
    "strconv"
)

// CreateTransaction membuat transaksi baru
func CreateTransaction(c *fiber.Ctx) error {
    // Get user dari token JWT
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(*config.JWTClaim)
    userID := claims.ID

    // Ambil data dari request
    var request struct {
        AlamatPengiriman uint `json:"alamat_pengiriman"`
        Produk           []struct {
            IDProduk  uint `json:"id_produk"`
            Kuantitas int  `json:"kuantitas"`
        } `json:"produk"`
        MethodBayar string `json:"method_bayar"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Format request tidak valid",
            "error":   err.Error(),
        })
    }

    // Hitung total harga dan validasi stok
    var totalHarga int
    var logProduks []models.LogProduk

    for _, item := range request.Produk {
        // Ambil data produk
        var product models.Product
        if err := database.DB.First(&product, item.IDProduk).Error; err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Produk tidak ditemukan",
            })
        }

        // Cek stok
        if product.Stok < item.Kuantitas {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Stok produk tidak mencukupi",
            })
        }

        // Konversi harga dari string ke int
        hargaKonsumen, err := strconv.Atoi(product.HargaKonsumen)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Format harga tidak valid",
                "error":   err.Error(),
            })
        }

        // Buat log produk
        logProduk := models.LogProduk{
            IDProduk:      product.ID,
            NamaProduk:    product.NamaProduk,
            Slug:          product.Slug,
            HargaReseller: product.HargaReseller,
            HargaKonsumen: product.HargaKonsumen,
            Deskripsi:     product.Deskripsi,
            IDToko:        product.IDToko,
            IDCategory:    product.IDCategory,
            CreatedAt:     time.Now(),
            UpdatedAt:     time.Now(),
        }

        if err := database.DB.Create(&logProduk).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Gagal membuat log produk",
                "error":   err.Error(),
            })
        }

        logProduks = append(logProduks, logProduk)
        totalHarga += item.Kuantitas * hargaKonsumen

        // Kurangi stok produk
        product.Stok -= item.Kuantitas
        if err := database.DB.Save(&product).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Gagal mengupdate stok produk",
                "error":   err.Error(),
            })
        }
    }

    // Buat transaksi
    trx := models.Trx{
        IDUser:           userID,
        AlamatPengiriman: request.AlamatPengiriman,
        HargaTotal:       totalHarga,
        KodeInvoice:      "INV-" + time.Now().Format("20060102150405"),
        MethodBayar:      request.MethodBayar,
        CreatedAt:        time.Now(),
        UpdatedAt:        time.Now(),
    }

    if err := database.DB.Create(&trx).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal membuat transaksi",
            "error":   err.Error(),
        })
    }

    // Buat detail transaksi
    for i, item := range request.Produk {
        // Konversi harga dari string ke int untuk detail transaksi
        hargaKonsumen, err := strconv.Atoi(logProduks[i].HargaKonsumen)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Format harga tidak valid",
                "error":   err.Error(),
            })
        }

        detail := models.DetailTrx{
            IDTrx:       trx.ID,
            IDLogProduk: logProduks[i].ID,
            IDToko:      logProduks[i].IDToko,
            Kuantitas:   item.Kuantitas,
            HargaTotal:  item.Kuantitas * hargaKonsumen,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }

        if err := database.DB.Create(&detail).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Gagal membuat detail transaksi",
                "error":   err.Error(),
            })
        }
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Transaksi berhasil dibuat",
    })
} 