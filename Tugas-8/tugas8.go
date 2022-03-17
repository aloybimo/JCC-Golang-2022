package main

import (
	"fmt"
	"math"
	"strconv"
	//"strings"
)

type segitigaSamaSisi struct{
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas*s.tinggi/2
}

func (s segitigaSamaSisi) keliling() int {
	return 3*s.alas
}
  
type persegiPanjang struct{
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang*p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2*(p.panjang+p.lebar)
}

type tabung struct{
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi*t.jariJari*t.jariJari*t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	var a = 2*math.Pi*t.jariJari*t.jariJari
	var b = 2*math.Pi*t.jariJari*t.tinggi
	return a+b
}

type balok struct{
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang*b.lebar*b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	var a = b.panjang*b.tinggi
	var c = b.lebar*b.tinggi
	var d = b.panjang*b.lebar
	return float64(2*(a+c+d))
}

type hitungbangunDatar interface{
	luas() int
	keliling() int
}

type hitungBangunRuang interface{
	volume() float64
	luasPermukaan() float64
}

type phone struct{
	name, brand string
	year int
	colors []string
}

func (p phone) tampilData(nama string) string{
	a := ""
	if nama == "name" {
		a += "name : " + p.name 
	} else if  nama == "brand" {
		a += "brand : "+p.brand 
	} else if  nama == "year" {
		a += "year : "+ strconv.Itoa(p.year)
	} else if  nama == "colors" {
		a += "colors : "
		for _, item := range p.colors{
			a += item + ", "
		}
	}
	fmt.Println(a)
	return a
}

type data interface {
	tampilData(string) string
}

func luasPersegi(sisi int, kondisi bool) {

}

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var bangunDatar hitungbangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{2, 3}
	fmt.Println("Luas Segitiga:",bangunDatar.luas())
	fmt.Println("Keliling Segitiga:",bangunDatar.keliling())

	bangunDatar = persegiPanjang{2, 3}
	fmt.Println("Luas persegi panjang:",bangunDatar.luas())
	fmt.Println("Keliling persegi panjang:",bangunDatar.keliling())

	bangunRuang = tabung{2, 3}
	fmt.Println("Volume tabung:",bangunRuang.volume())
	fmt.Println("Luas permukaan tabung:",bangunRuang.luasPermukaan())

	bangunRuang = balok{2, 3, 4}
	fmt.Println("Volume balok:",bangunRuang.volume())
	fmt.Println("Luas permukaan balok:",bangunRuang.luasPermukaan())

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var galaxy20 data
	galaxy20 = phone{"Samsung Galaxy Note 20", "Samsung", 2020, []string{"Mystic Bronze","Mystic White","Mystic Black"}}
	galaxy20.tampilData("name")
	galaxy20.tampilData("brand")
	galaxy20.tampilData("year")
	galaxy20.tampilData("colors")

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}
	// tulis jawaban anda disini
	var kalimatPertama = prefix.(string)
	var kalimatKedua = ""
	for _, value := range kumpulanAngkaPertama.([]int) {
		kalimatKedua += strconv.Itoa(value) + " + "
	}
	var kalimatKetiga = ""
	for i, value := range kumpulanAngkaKedua.([]int) {
		if i != len(kumpulanAngkaKedua.([]int))-1 {
			kalimatKetiga += strconv.Itoa(value) + " + "	
		} else {
			kalimatKetiga += strconv.Itoa(value) + " = "
		} 
	}
	//var kalimatKedua = strings.Join(kumpulanAngkaPertama.([]int), ", ")
	//var kalimatKetiga = strings.Join(kumpulanAngkaKedua.([]int), ", ")
	a := 0
	for _, item := range kumpulanAngkaPertama.([]int){
		a += item
	}
	for _, item := range kumpulanAngkaKedua.([]int){
		a += item
	}
	fmt.Println(kalimatPertama + kalimatKedua + kalimatKetiga + strconv.Itoa(a))

}