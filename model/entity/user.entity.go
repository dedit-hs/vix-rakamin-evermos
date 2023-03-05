package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Nama string `json:"name"`
	KataSandi string `json:"kataSandi"`
	Telepon string `json:"telepon"`
	TanggalLahir time.Time `json:"tanggalLahir"`
	JenisKelamin string `json:"jenisKelamin"`
	Tentang string `json:"tentang"`
	Pekerjaan string `json:"pekerjaan"`
	Email string `json:"email"`
	IDProvinsi uint8 `json:"idProvinsi"`
	IDKota uint8 `json:"idKota"`
	IsAdmin bool `json:"isAdmin"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}