package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Nama         string     `json:"nama"`
	KataSandi    string     `json:"kata_sandi"`
	Notelp       string     `gorm:"unique" json:"notelp"`
	TanggalLahir *time.Time `json:"tanggal_lahir"`
	JenisKelamin string     `json:"jenis_kelamin"`
	Tentang      string     `json:"tentang"`
	Pekerjaan    string     `json:"pekerjaan"`
	Email        string     `json:"email"`
	IDProvinsi   string     `json:"id_provinsi"`
	IDKota       string     `json:"id_kota"`
	IsAdmin      bool       `json:"is_admin"`
	UpdatedAt    time.Time  `json:"updated_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

func (User) TableName() string {
    return "user"
}
