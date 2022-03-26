package models

import (
  "time"
)

type (
  // Nilai
  NilaiMahasiswa struct {
    ID        	int       `json:"id"`
    IndeksNilai	string	`json:"indeksNilai"`
	  Skor		int		`json:"skor"`
    MahasiswaID 		int	`json:"mahasiswa_id"`
	  MataKuliahID	int	`json:"mata_kuliah_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
  }
)