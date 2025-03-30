package models

import "time"

type Category struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    NamaCategory string    `json:"nama_category" gorm:"type:varchar(255)"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
} 