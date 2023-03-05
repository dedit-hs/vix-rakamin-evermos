package request

import "time"

type UserCreateRequest struct {
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
}