package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func luasPersegiPanjang(panjang int, lebar int) string{
	return strconv.Itoa(panjang*lebar)
}

func kelilingPersegiPanjang(panjang int, lebar int) string{
	return strconv.Itoa(2*(panjang+lebar))
}

func volumeBalok(panjang int, lebar int, tinggi int) string{
	return strconv.Itoa(panjang*lebar*tinggi)
}

func introduce(nama string, gender string, job string, age string) (kalimat string) {
	if gender == "laki-laki" {
		kalimat = "Pak "+nama+" adalah seorang "+job+" yang berusia "+age+" tahun"
	} else if gender == "perempuan" {
		kalimat = "Bu "+nama+" adalah seorang "+job+" yang berusia "+age+" tahun"
	}
	return
} 

func buahFavorit(nama string, buah ...string) (kalimat string) {
	kalimatBuah := ""
	for i, isi := range buah {
		if i < len(buah)-1 {
			kalimatBuah = kalimatBuah +`"`+isi+`"`+", "
		} else {
			kalimatBuah = kalimatBuah +`"`+isi+`"`
		}
	kalimat = "halo nama saya "+nama+" dan buah favorit saya adalah "+kalimatBuah
	}
	return
}

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	panjang := 12
	lebar := 4
	tinggi := 8
  	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)
	fmt.Println(luas) 
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("john", buah...)
	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	
	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var kontenFilm = func(title string, jam string, genre string, tahun string) {
		var konten  = map[string]string{}
		konten["genre"] = genre
		konten["jam"] = jam
		konten["tahun"] = tahun
		konten["title"] = title
		dataFilm = append(dataFilm, konten)
	}

	tambahDataFilm := kontenFilm
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
	for _, item := range dataFilm {
		fmt.Println(item)
	}
}