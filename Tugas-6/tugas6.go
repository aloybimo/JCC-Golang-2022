package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func luaslingkaran(jarijari float64, luasLingkaran *float64) {
	*luasLingkaran = 3.14 * jarijari * jarijari
}

func kelilinglingkaran(jarijari float64, kelilingLingkaran *float64) {
	*kelilingLingkaran = 3.14 * 2 * jarijari
}

func introduce(sentence *string, name string, gender string, job string, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak "+name+" adalah seorang "+job+" yang berusia "+age+" tahun"
	} else if gender == "perempuan" {
		*sentence = "Bu "+name+" adalah seorang "+job+" yang berusia "+age+" tahun"
	}
}

func tambahBuah (buah *[]string, jenisbuah ...string) {
	*buah = append(*buah, jenisbuah...)
}

func tambahDataFilm(title string, jam string, genre string, tahun string, dataFilm *[]map[string]string) {
	var konten  = map[string]string{}
	konten["genre"] = genre
	konten["duration"] = jam
	konten["year"] = tahun
	konten["title"] = title
	*dataFilm = append(*dataFilm, konten)
}

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var luasLingkaran float64 
	var kelilingLingkaran float64
	fmt.Println("Luas Lingkaran: ", luasLingkaran)
	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran)
	luaslingkaran(5.00, &luasLingkaran)
	kelilinglingkaran(5.00, &kelilingLingkaran)
	fmt.Println("Luas Lingkaran: ", luasLingkaran)
	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran)	

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	var buah = []string{}
	tambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	for i, isi := range buah {
		fmt.Println(strconv.Itoa(i+1)+". "+isi)
	}
	
	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	for i, item := range dataFilm {
		if i != 2 {
			fmt.Println(strconv.Itoa(i+1)+". title : "+item["title"])
			fmt.Println("   duration : "+item["duration"])
			fmt.Println("   genre : "+item["genre"])
			fmt.Println("   year : "+item["year"])
			fmt.Println("")
		} else {
			fmt.Println(strconv.Itoa(i+1)+". genre : "+item["genre"])
			fmt.Println("   year : "+item["year"])
			fmt.Println("   title : "+item["title"])
			fmt.Println("   duration : "+item["duration"])
			fmt.Println("")
		}
		
	}
}