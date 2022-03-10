package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"
	newpanjangPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	newlebarPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
	newalasSegitiga, _ := strconv.Atoi(alasSegitiga)
	newtinggiSegitiga, _ := strconv.Atoi(tinggiSegitiga)
	var kelilingPersegiPanjang int = 2*(newpanjangPersegiPanjang + newlebarPersegiPanjang)
	var luasSegitiga int = newalasSegitiga * newtinggiSegitiga/2
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var nilaiJohn = 80
	var nilaiDoe = 50
	if nilaiJohn >= 80 {
		fmt.Println("Indexnya John A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Indeksnya John B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Indeksnya John C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Indeksnya John D")
	} else {
		fmt.Println("Indeksnya John E")
	}
	if nilaiDoe >= 80 {
		fmt.Println("Indexnya John A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Indeksnya Doe B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Indeksnya Doe C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Indeksnya Doe D")
	} else {
		fmt.Println("Indeksnya Doe E")
	}

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	var tanggal = 17;
	var bulan = 8; 
	var tahun = 1945;
	tanggal = 16;
	bulan = 6;
	tahun = 2000;
	stringTanggal := strconv.Itoa(tanggal)
	stringBulan := strconv.Itoa(bulan)
	stringTahun := strconv.Itoa(tahun)
	switch stringBulan{
	case "6":
		fmt.Println(stringTanggal+" Juni "+stringTahun)
	default:
		fmt.Println(stringTanggal+" "+stringBulan+" "+stringTahun)
	}

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var tahunLahir = 2000
	if tahunLahir>= 1944 && tahunLahir<= 1964 {
		fmt.Println("Baby boomer, kelahiran 1944 s.d 1964")
	} else if tahunLahir>= 1965 && tahunLahir<= 1979 {
		fmt.Println("Generasi X, , kelahiran 1965 s.d 1979")
	} else if tahunLahir>= 1980 && tahunLahir<= 1994 {
		fmt.Println("Generasi Y (Millenials), kelahiran 1980 s.d 1994")
	} else if tahunLahir>= 1995 && tahunLahir<= 2015 {
		fmt.Println("Generasi Z, , kelahiran 1995 s.d 2015")
	}
	
}