package main

import (
	"fmt"
	"strconv"
	"flag"
	"errors"
	"math"
	"time"
	
)

func generateKalimat(kalimat string, tahun int) {
	fmt.Println("Jawaban Soal 1")
	fmt.Println(kalimat, strconv.Itoa(tahun))
	fmt.Println("")
}

func showError(){
	message := recover()
	fmt.Println("Terjadi Error: ", message)
}

func kelilingSegitigaSamaSisi(sisi int, condition bool)(string, error){
	if condition == true && sisi != 0{
		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(sisi*3) + " cm", nil
	} else if condition == false && sisi != 0{
		return strconv.Itoa(sisi*3), nil
	} else if condition == true && sisi == 0{
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if condition == false && sisi == 0{
		defer showError()
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
	return "", errors.New("a")
}

func cetakAngka(angka *int) {
	fmt.Println("")
	fmt.Println("Jawaban Soal 3: ")
	fmt.Println(*angka)
	fmt.Println("")
}

func tambahAngka(angka int, simpanan *int) {
	*simpanan += angka
}

func tambahData(tempat *[]string, data ...string) {
	*tempat = append(*tempat, data...)
}

func luasLingkaran(jariJari int) {
	luas := math.Pi * math.Pow(float64(jariJari),2)
	fmt.Println(math.Round(luas))
}

func kelilingLingkaran(jariJari int) {
	keliling := math.Pi * 2 * float64(jariJari)
	fmt.Println(math.Round(keliling))
}

func luasPersegiPanjang(panjang, lebar int) {
	luas := panjang*lebar
	fmt.Println(luas)
}

func kelilingPersegiPanjang(panjang, lebar int) {
	keliling := 2 * (panjang + lebar)
	fmt.Println(keliling)
}

func main() {
	// [SOAL 3]deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	//soal 1
	defer generateKalimat("Golang Backend Development", 2021)

	//soal 2 (BELUM SELESAI)
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//soal 3
	fmt.Println(" ")
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var phones = []string{}
	tambahData(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	for i, item := range phones {
		fmt.Println(strconv.Itoa(i+1)+". "+item)
		time.Sleep(time.Second)
	}

	//soal 5
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 5")
	luasLingkaran(7)
	kelilingLingkaran(7)
	luasLingkaran(10)
	kelilingLingkaran(10)
	luasLingkaran(15)
	kelilingLingkaran(15)

	//soal 6
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 6")
	var panjang = flag.Int("panjang", 6, "Panjang persegi panjang")
	var lebar = flag.Int("lebar", 8, "Lebar persegi panjang")
	flag.Parse()
	luasPersegiPanjang(*panjang, *lebar)
	kelilingPersegiPanjang(*panjang, *lebar)

}