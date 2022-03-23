package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct{
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID uint
}

type Temporary struct{
	Nama string `json:"Nama"`
	MataKuliah string `json:"MataKuliah"`
	Nilai uint `json:"Nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

//Authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
	  	if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
	    }
  
	  	if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
	 	}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// GetNilai
func getNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// postNilai
func postNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nilaiBaru NilaiMahasiswa
	var simpanSementara Temporary
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&simpanSementara); err != nil {
				log.Fatal(err)
			}
			var indeks string
			if simpanSementara.Nilai >= 80 && simpanSementara.Nilai <= 100 {
				indeks = "A"
			} else if simpanSementara.Nilai >= 70 && simpanSementara.Nilai < 80 {
				indeks = "B"
			} else if simpanSementara.Nilai >= 60 && simpanSementara.Nilai < 70 {
				indeks = "C"
			} else if simpanSementara.Nilai >= 50 && simpanSementara.Nilai < 60 {
				indeks = "D"
			} else if simpanSementara.Nilai < 50 {
				indeks = "E"
			} else {
				indeks = "-"
			}
			nilaiBaru = NilaiMahasiswa{
				Nama: simpanSementara.Nama,
				MataKuliah: simpanSementara.MataKuliah,
				IndeksNilai: indeks,
				Nilai: simpanSementara.Nilai,
				ID: uint(len(nilaiNilaiMahasiswa) + 1),
			}
		} else {
			// parse dari form
			nama := r.PostFormValue("Nama")
			matakuliah := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			nilai, _ := strconv.Atoi(getNilai)
			var indeks string
			if nilai >= 80 && nilai <= 100 {
				indeks = "A"
			} else if nilai >= 70 && nilai < 80 {
				indeks = "B"
			} else if nilai >= 60 && nilai < 70 {
				indeks = "C"
			} else if nilai >= 50 && nilai < 60 {
				indeks = "D"
			} else if nilai < 50 {
				indeks = "E"
			} else {
				indeks = "-"
			}
			nilaiBaru = NilaiMahasiswa{
				Nama: nama,
				MataKuliah: matakuliah,
				IndeksNilai: indeks,
				Nilai: uint(nilai),
				ID: uint(len(nilaiNilaiMahasiswa) + 1),
			}
		}

		dataNilai, _ := json.Marshal(nilaiBaru) // to byte
		w.Write(dataNilai)                // cetak di browser
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiBaru)
		return
	} 
	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}


func main() {
	http.HandleFunc("/nilai", getNilai)
	http.Handle("/post_nilai", Auth(http.HandlerFunc(postNilai)))
	fmt.Println("server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}