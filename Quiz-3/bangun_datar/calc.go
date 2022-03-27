package bangun_datar

import (
	"fmt"
	"math"
)

//Segitiga Sama Sisi
func LuasSegitiga(ch chan string, alas, tinggi int) {
	luas := alas*tinggi/2
	str := fmt.Sprintf("%d", luas)
	ch <- str
}

func KelilingSegitiga(ch chan string, alas int) {
	keliling := 3*alas
	str := fmt.Sprintf("%d", keliling)
	ch <- str
}

//Persegi
func LuasPersegi(ch chan string, sisi int) {
	luas := sisi*sisi
	str := fmt.Sprintf("%d", luas)
	ch <- str
}

func KelilingPersegi(ch chan string, sisi int) {
	keliling := 4*sisi
	str := fmt.Sprintf("%d", keliling)
	ch <- str
}

//Persegi Panjang
func LuasPersegiPanjang(ch chan string, panjang, lebar int) {
	luas := panjang*lebar
	str := fmt.Sprintf("%d", luas)
	ch <- str
}

func KelilingPersegiPanjang(ch chan string, panjang, lebar int) {
	keliling := 2*(panjang+lebar)
	str := fmt.Sprintf("%d", keliling)
	ch <- str
}

//Lingkaran
func LuasLingkaran(ch chan string, jariJari int) {
	luas := math.Pi * math.Pow(float64(jariJari),2)
	str := fmt.Sprintf("%f", luas)
	ch <- str
}

func KelilingLingkaran(ch chan string, jariJari int) {
	keliling := math.Pi * 2 * float64(jariJari)
	str := fmt.Sprintf("%f", keliling)
	ch <- str
}

//Jajar Genjang
func LuasJajarGenjang(ch chan string, alas, tinggi int) {
	luas := alas*tinggi
	str := fmt.Sprintf("%d", luas)
	ch <- str
}

func KelilingJajarGenjang(ch chan string, alas, sisi int) {
	keliling := 2*(alas+sisi)
	str := fmt.Sprintf("%d", keliling)
	ch <- str
}
