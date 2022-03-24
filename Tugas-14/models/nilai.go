package models

import (
  "time"
)

type (
  // Nilai
  NilaiMahasiswa struct {
    ID        	int       `json:"id"`
	  Nama 		string	`json:"nama"`
	  MataKuliah	string	`json:"mataKuliah"`
    IndeksNilai	string	`json:"indeksNilai"`
	  Nilai		int		`json:"nilai"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
  }
)