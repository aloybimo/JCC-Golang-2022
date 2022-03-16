package main

import (
	"fmt"
	"strconv"
	//"strings"
)

type buah struct {
	nama string
	warna string
	adaBijinya bool
	harga int
}

type segitiga struct{
	alas, tinggi int
}
  
type persegi struct{
	sisi int
}
  
type persegiPanjang struct{
	panjang, lebar int
}

type phone struct{
	name, brand string
	year int
	colors []string
}

type movie struct{
	title, duration, genre string
	year int
}

func (bidang segitiga) luasSegitiga() {
	fmt.Println("Luas Segitiga :",bidang.alas * bidang.tinggi / 2)
}

func (bidang persegi) luasPersegi() {
	fmt.Println("Luas Persegi :",bidang.sisi * bidang.sisi)
}

func (bidang persegiPanjang) luasPersegiPanjang() {
	fmt.Println("Luas Persegi Panjang :",bidang.panjang * bidang.lebar)
}

func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
} 

func tambahDataFilm(title string, menit int, genre string, tahun int, m *[]movie) {
	var konten  = movie{}
	konten.genre = genre
	konten.duration = strconv.Itoa(menit/60)+" jam"
	konten.year = tahun
	konten.title = title
	*m = append(*m, konten)
}

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var nanas buah
	nanas.nama = "Nanas"
	nanas.warna = "Kuning"
	nanas.adaBijinya = false
	nanas.harga = 9000

	var jeruk buah
	jeruk.nama = "Jeruk"
	jeruk.warna = "Oranye"
	jeruk.adaBijinya = true
	jeruk.harga = 8000

	var semangka buah
	semangka.nama = "Semangka"
	semangka.warna = "Hijau & Merah"
	semangka.adaBijinya = true
	semangka.harga = 10000

	var pisang buah
	pisang.nama = "Pisang"
	pisang.warna = "Kuning"
	pisang.adaBijinya = false
	pisang.harga = 5000

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var bidang1 = segitiga{2, 3}
	var bidang2 = persegi{3}
	var bidang3 = persegiPanjang{2,3}
	bidang1.luasSegitiga()
	bidang2.luasPersegi()
	bidang3.luasPersegiPanjang()

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	var redmi11 = phone{"Redmi 11","Xiaomi",2021, []string{"Biru"}}
	fmt.Println(redmi11)
	redmi11.tambahWarna("Hitam")
	fmt.Println("---------Tambah Warna---------")
	fmt.Println(redmi11)

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	for i, item := range dataFilm {
		if i != 2 {
			fmt.Println(strconv.Itoa(i+1)+". title : "+item.title)
			fmt.Println("   duration : "+item.duration)
			fmt.Println("   genre : "+item.genre)
			fmt.Println("   year :",item.year)
			fmt.Println("")
		} else {
			fmt.Println(strconv.Itoa(i+1)+". genre : "+item.genre)
			fmt.Println("   year :",item.year)
			fmt.Println("   title : "+item.title)
			fmt.Println("   duration : "+item.duration)
			fmt.Println("")
		}
	}
}