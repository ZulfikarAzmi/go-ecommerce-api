package models

import "time"

type Toko struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDUser    uint      `json:"id_user" gorm:"not null"`
	NamaToko  string    `json:"nama_toko" gorm:"size:255"`
	URLFoto   string    `json:"url_foto" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
} 