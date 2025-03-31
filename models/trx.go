package models

import "time"

type Trx struct {
    ID                uint      `json:"id" gorm:"primaryKey"`
    IDUser            uint      `json:"id_user"` // Foreign key untuk user
    AlamatPengiriman  uint      `json:"alamat_pengiriman"` // Foreign key untuk alamat
    HargaTotal        int       `json:"harga_total"`
    KodeInvoice       string    `json:"kode_invoice" gorm:"type:varchar(255)"`
    MethodBayar      string    `json:"method_bayar" gorm:"type:varchar(255)"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
    User              User      `json:"user" gorm:"foreignKey:IDUser"` // Relasi ke model User
    Alamat            Alamat    `json:"alamat" gorm:"foreignKey:AlamatPengiriman"` // Relasi ke model Alamat
}

// TableName menentukan nama tabel yang digunakan di database
func (Trx) TableName() string {
    return "trx"
} 