package main

import (
	"fmt"
	"sort"
	"time"
	"math"
	"sync"
	"strconv"
)

func tampilData(data []string, wg *sync.WaitGroup, sec int) {
	for i, item := range data {
		fmt.Println(i+1, ". ", item)
		time.Sleep(time.Second * time.Duration(sec))
	}
	wg.Done()
}

func getMovies(ch chan string, data ...string) {
	fmt.Println("List Movies: ")
	for i, isi := range data {
		kirim := strconv.Itoa(i+1)+". "+isi
		ch <- kirim
	}
	close(ch)
}	

func luasLingkaran(ch chan float64, jariJari int) {
	luas := math.Pi * math.Pow(float64(jariJari),2)
	ch <- luas
}

func kelilingLingkaran(ch chan float64, jariJari int) {
	keliling := math.Pi * 2 * float64(jariJari)
	ch <- keliling
}

func volumeTabung(ch chan float64, jariJari, tinggi int) {
	volume := math.Pi * math.Pow(float64(jariJari),2) * float64(tinggi)
	ch <- volume
}

func luasPersegiPanjang(ch chan int, panjang, lebar int) {
	luas := panjang*lebar
	ch <- luas
}

func kelilingPersegiPanjang(ch chan int, panjang, lebar int) {
	keliling := 2 * (panjang + lebar)
	ch <- keliling
}

func volumeBalok(ch chan int, panjang, lebar, tinggi int) {
	volume := panjang*lebar*tinggi
	ch <- volume
}

func main() {
	//soal 1
	fmt.Println("Jawaban Soal 1")
	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	wg.Add(1)
	go tampilData(phones, &wg, 1)
	wg.Wait()

	//soal 2
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 2")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}

	//soal 3
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 3")
	hasil := make(chan float64)
	
	go luasLingkaran(hasil, 8)
	terima := <-hasil
	fmt.Println(terima)
	
	go luasLingkaran(hasil, 14)
	terima = <-hasil
	fmt.Println(terima)

	go luasLingkaran(hasil, 20)
	terima = <-hasil
	fmt.Println(terima)
	
	go kelilingLingkaran(hasil, 8)
	terima = <-hasil
	fmt.Println(terima)
	
	go kelilingLingkaran(hasil,14)
	terima = <-hasil
	fmt.Println(terima)
	
	go kelilingLingkaran(hasil,20)
	terima = <-hasil
	fmt.Println(terima)
	
	go volumeTabung(hasil,8,10)
	terima = <-hasil
	fmt.Println(terima)
	
	go volumeTabung(hasil,14,10)
	terima = <-hasil
	fmt.Println(terima)
	
	go volumeTabung(hasil,20,10)
	terima = <-hasil
	fmt.Println(terima)

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	var chLuas = make(chan int)
	go luasPersegiPanjang(chLuas, 2,3)
	var chKeliling = make(chan int)
	go kelilingPersegiPanjang(chKeliling, 2,3)
	var chVolume = make(chan int)
	go volumeBalok(chVolume, 2,3,4)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-chLuas:
		  fmt.Println("Luas :"+strconv.Itoa(luas))
		case keliling := <-chKeliling:
		  fmt.Println("Keliling :"+strconv.Itoa(keliling))
		case volume := <-chVolume:
			fmt.Println("Volume :"+strconv.Itoa(volume))
	  	}
	}

}