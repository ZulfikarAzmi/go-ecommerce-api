package models

import "time"

type Alamat struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDUser       uint      `json:"id_user" gorm:"not null"`
	JudulAlamat  string    `json:"judul_alamat" gorm:"size:255;not null"`
	NamaPenerima string    `json:"nama_penerima" gorm:"size:255;not null"`
	NoTelp       string    `json:"no_telp" gorm:"size:255;not null"`
	DetailAlamat string    `json:"detail_alamat" gorm:"type:text;not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user" gorm:"foreignKey:IDUser"`
}

// Menentukan nama tabel
func (Alamat) TableName() string {
	return "alamat" // Pastikan nama tabel adalah "alamat"
} 