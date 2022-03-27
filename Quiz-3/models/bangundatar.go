package models

type (
	// Segitiga sama sisi
	Segitiga struct {
		Alas int
		Tinggi int
		Hitung string
	}

	Persegi struct {
		Sisi int
		Hitung string
	}

	PersegiPanjang struct {
		Panjang int
		Lebar int
		Hitung string
	}

	Lingkaran struct {
		JariJari int
		Hitung string
	}

	JajarGenjang struct {
		Sisi int
		Alas int
		Tinggi int
		Hitung string
	}
  )