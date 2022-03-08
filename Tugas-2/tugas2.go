package main

import "fmt"
import "strings"
import "strconv"

func main() {
	//soal 1
	var kata1 = "Jabar"
	var kata2 = "Coding"
	var kata3 = "Camp"
	var kata4 = "Golang"
	var kata5 = "2022"
	fmt.Println(`"` + kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5 + `"`)

	//soal 2
	halo := "Halo Dunia"
	remove := "Dunia"
	newWord := "Golang"
	newHalo := strings.Replace(halo, remove, newWord, 1)
	fmt.Println(newHalo)

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)
	fmt.Println(kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat)

	//soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";
	newangkaPertama, _ := strconv.Atoi(angkaPertama)
	newangkaKedua, _ := strconv.Atoi(angkaKedua)
	newangkaKetiga, _ := strconv.Atoi(angkaKetiga)
	newangkaKeempat, _ := strconv.Atoi(angkaKeempat)
	var jumlah = newangkaPertama + newangkaKedua + newangkaKetiga + newangkaKeempat
	fmt.Println(jumlah)

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2022
	kalimat = strings.Replace(kalimat, "halo", "Hi", 2)
	newangka := strconv.Itoa(angka)
	fmt.Println(`"`+kalimat + `" - ` + newangka)
}