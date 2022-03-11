package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	for i := 1; i < 21; i++ {
		switch {
		case i%2 == 1 && i%3 == 0:
			fmt.Println(strconv.Itoa(i)," - I Love Coding")
		case i%2 == 1 && i%3 != 0:
			fmt.Println(strconv.Itoa(i)," - JCC")
		default:
			fmt.Println(strconv.Itoa(i)," - Candradimuka")
		}
	}

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var pagar = "#"
	for i := 1; i<8; i++ {
		fmt.Println(strings.Repeat(pagar, i))
	}

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	kalimatBaru := ""
	for i:=2; i<len(kalimat); i++ {
		if len(kalimatBaru)==0 {
			kalimatBaru += kalimat[i]
		} else {
			kalimatBaru = kalimatBaru + " " + kalimat[i]
		}
	}
	fmt.Println(kalimatBaru)
	
	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, sayur := range sayuran {
		fmt.Println(strconv.Itoa(i+1)+". "+sayur)
	}

	//soal 5
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 5")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volume = 1
	for key, element := range satuan {
		fmt.Println(key,"=",element)
		volume *= element
	}
	fmt.Println("Volume Balok =",volume)

}