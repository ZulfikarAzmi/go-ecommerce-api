package models

import (
	"time"
)

type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Nama         string     `json:"nama"`
	KataSandi    string     `json:"kata_sandi"`
	Notelp       string     `gorm:"unique" json:"notelp"`
	TanggalLahir *time.Time `gorm:"type:date" json:"tanggal_lahir"`
	JenisKelamin string     `json:"jenis_kelamin"`
	Tentang      string     `json:"tentang"`
	Pekerjaan    string     `json:"pekerjaan"`
	Email        string     `json:"email"`
	IDProvinsi   string     `json:"id_provinsi"`
	IDKota       string     `json:"id_kota"`
	IsAdmin      bool       `json:"is_admin"`
	UpdatedAt    time.Time  `gorm:"type:date" json:"updated_at"`
	CreatedAt    time.Time  `gorm:"type:date" json:"created_at"`
}

func (User) TableName() string {
	return "user"
}
