package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-ecommerce-api/models"
)

var DB *gorm.DB

func Connect() {
	// Konfigurasi database
	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	dbname := "ecommerce-api-go"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	// Membuat koneksi ke database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	fmt.Println("Koneksi Database Berhasil!")

	// Auto Migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Toko{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: ", err)
	}
}
