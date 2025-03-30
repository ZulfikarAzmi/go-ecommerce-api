package models

import (
	"time"
	"github.com/gosimple/slug"
)

type Product struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	NamaProduk    string    `json:"nama_produk" gorm:"type:varchar(255)"`
	Slug          string    `json:"slug" gorm:"type:varchar(255);unique"`
	HargaReseller string    `json:"harga_reseller" gorm:"type:varchar(255)"`
	HargaKonsumen string    `json:"harga_konsumen" gorm:"type:varchar(255)"`
	Stok          int       `json:"stok" gorm:"default:0"`
	Deskripsi     string    `json:"deskripsi" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IDToko        uint      `json:"id_toko"`
	IDCategory    uint      `json:"id_category"`
	Toko          Toko      `json:"toko" gorm:"foreignKey:IDToko"`
	Category      Category  `json:"category" gorm:"foreignKey:IDCategory"`
}

// BeforeCreate adalah hook GORM yang akan dijalankan sebelum record dibuat
func (p *Product) BeforeCreate() error {
	p.Slug = slug.Make(p.NamaProduk)
	return nil
} 