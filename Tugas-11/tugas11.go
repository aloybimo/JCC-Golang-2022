package main

import (
	"fmt"
	"sort"
	"time"
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
	

	//soal 4
	fmt.Println(" ")
	fmt.Println("Jawaban Soal 4")
	

}