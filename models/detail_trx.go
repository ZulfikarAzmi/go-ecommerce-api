package models

import "time"

type DetailTrx struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    IDTrx       uint      `json:"id_trx"`
    IDLogProduk  uint      `json:"id_log_produk"`
    IDToko      uint      `json:"id_toko"`
    Kuantitas   int       `json:"kuantitas"`
    HargaTotal  int       `json:"harga_total"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Trx         Trx       `json:"trx" gorm:"foreignKey:IDTrx"`
    LogProduk   LogProduk  `json:"log_produk" gorm:"foreignKey:IDLogProduk"`
    Toko        Toko      `json:"toko" gorm:"foreignKey:IDToko"`
}

// TableName menentukan nama tabel yang digunakan di database
func (DetailTrx) TableName() string {
    return "detail_trx"
} 