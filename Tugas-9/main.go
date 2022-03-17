package main

import (
	. "Tugas-9/library"
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var bangunDatar HitungbangunDatar
	var bangunRuang HitungBangunRuang

	bangunDatar = SegitigaSamaSisi{2, 3}
	fmt.Println("Luas Segitiga:",bangunDatar.Luas())
	fmt.Println("Keliling Segitiga:",bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{2, 3}
	fmt.Println("Luas persegi panjang:",bangunDatar.Luas())
	fmt.Println("Keliling persegi panjang:",bangunDatar.Keliling())

	bangunRuang = Tabung{2, 3}
	fmt.Println("Volume tabung:",bangunRuang.Volume())
	fmt.Println("Luas permukaan tabung:",bangunRuang.LuasPermukaan())

	bangunRuang = Balok{2, 3, 4}
	fmt.Println("Volume balok:",bangunRuang.Volume())
	fmt.Println("Luas permukaan balok:",bangunRuang.LuasPermukaan())

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var galaxy20 Data
	galaxy20 = Phone{"Samsung Galaxy Note 20", "Samsung", 2020, []string{"Mystic Bronze","Mystic White","Mystic Black"}}
	galaxy20.TampilData("name")
	galaxy20.TampilData("brand")
	galaxy20.TampilData("year")
	galaxy20.TampilData("colors")

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))
	

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