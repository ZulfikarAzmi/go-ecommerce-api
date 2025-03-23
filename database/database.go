package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	"go-ecommerce-api/models"
)

var DB *gorm.DB

func Connect() {
	// Konfigurasi koneksi database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Jika variabel environment tidak diatur, gunakan nilai default
	if dbUser == "" {
		dbUser = "root" // Default XAMPP username
	}
	if dbPass == "" {
		dbPass = "" // Default XAMPP password (kosong)
	}
	if dbHost == "" {
		dbHost = "127.0.0.1:3306" // Default XAMPP host
	}
	if dbName == "" {
		dbName = "ecommerce-api-go" // Nama database Anda
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)

	// Membuka koneksi ke database MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// AutoMigrate untuk membuat tabel User
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected and migrated successfully!")
}
