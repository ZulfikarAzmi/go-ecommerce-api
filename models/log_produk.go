package models

import "time"

type LogProduk struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    IDProduk      uint      `json:"id_produk"`
    NamaProduk    string    `json:"nama_produk" gorm:"type:varchar(255)"`
    Slug          string    `json:"slug" gorm:"type:varchar(255)"`
    HargaReseller string    `json:"harga_reseller" gorm:"type:varchar(255)"`
    HargaKonsumen string    `json:"harga_konsumen" gorm:"type:varchar(255)"`
    Deskripsi     string    `json:"deskripsi" gorm:"type:text"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    IDToko        uint      `json:"id_toko"`
    IDCategory    uint      `json:"id_category"`
    Toko          Toko      `json:"toko" gorm:"foreignKey:IDToko"`
    Category      Category  `json:"category" gorm:"foreignKey:IDCategory"`
}

// TableName menentukan nama tabel yang digunakan di database
func (LogProduk) TableName() string {
    return "log_produk"
} 