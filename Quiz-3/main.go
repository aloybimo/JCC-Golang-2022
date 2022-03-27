package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"context"
	"encoding/json"
	"quiz3/bangun_datar"
	"quiz3/models"
	"quiz3/categories"
	"quiz3/books"
	"quiz3/utils"
	"github.com/julienschmidt/httprouter"
)

var hasilHitung = make(chan string)

//BASIC AUTH
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
	  	if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
	    }
  
	  	if uname == "admin" && pwd == "password" || uname == "editor" && pwd == "secret" || uname == "trainer" && pwd == "rahasia" {
			next(w, r, ps)
			return
	 	}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	}
}

//BANGUN DATAR
//GET Segitiga Sama Sisi
func GetSegitiga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var segitiga models.Segitiga
	segitiga.Alas, _ = strconv.Atoi(r.URL.Query().Get("alas"))
	segitiga.Tinggi, _ = strconv.Atoi(r.URL.Query().Get("tinggi"))
	segitiga.Hitung = r.URL.Query().Get("hitung")
	if segitiga.Hitung == "keliling" {
		go bangun_datar.KelilingSegitiga(hasilHitung, segitiga.Alas)
	} else if segitiga.Hitung == "luas" {
		go bangun_datar.LuasSegitiga(hasilHitung, segitiga.Alas, segitiga.Tinggi)
	}
	a := <- hasilHitung
	kalimat := "Hasil Perhitungan: " + a
	utils.ResponseJSON(w, kalimat, http.StatusOK)
}

//GET Persegi
func GetPersegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var persegi models.Persegi
	persegi.Sisi, _ = strconv.Atoi(r.URL.Query().Get("sisi"))
	persegi.Hitung = r.URL.Query().Get("hitung")
	if persegi.Hitung == "keliling" {
		go bangun_datar.KelilingPersegi(hasilHitung, persegi.Sisi)
	} else if persegi.Hitung == "luas" {
		go bangun_datar.LuasPersegi(hasilHitung, persegi.Sisi)
	} 
	a := <- hasilHitung
	kalimat := "Hasil Perhitungan: " + a
	utils.ResponseJSON(w, kalimat, http.StatusOK)
}

//GET Persegi Panjang
func GetPersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var persegiPanjang models.PersegiPanjang
	persegiPanjang.Panjang, _ = strconv.Atoi(r.URL.Query().Get("panjang"))
	persegiPanjang.Lebar, _ = strconv.Atoi(r.URL.Query().Get("lebar"))
	persegiPanjang.Hitung = r.URL.Query().Get("hitung")
	if persegiPanjang.Hitung == "keliling" {
		go bangun_datar.KelilingPersegiPanjang(hasilHitung, persegiPanjang.Panjang, persegiPanjang.Lebar)
	} else if persegiPanjang.Hitung == "luas" {
		go bangun_datar.LuasPersegiPanjang(hasilHitung, persegiPanjang.Panjang, persegiPanjang.Lebar)
	} 
	a := <- hasilHitung
	kalimat := "Hasil Perhitungan: " + a
	utils.ResponseJSON(w, kalimat, http.StatusOK)
}

//GET Lingkaran
func GetLingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var lingkaran models.Lingkaran
	lingkaran.JariJari, _ = strconv.Atoi(r.URL.Query().Get("jariJari"))
	lingkaran.Hitung = r.URL.Query().Get("hitung")
	if lingkaran.Hitung == "keliling" {
		go bangun_datar.KelilingLingkaran(hasilHitung, lingkaran.JariJari)
	} else if lingkaran.Hitung == "luas" {
		go bangun_datar.LuasLingkaran(hasilHitung, lingkaran.JariJari)
	} 
	a := <- hasilHitung
	kalimat := "Hasil Perhitungan: " + a
	utils.ResponseJSON(w, kalimat, http.StatusOK)
}

//GET Jajar Genjang
func GetJajarGenjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var jajarGenjang models.JajarGenjang
	jajarGenjang.Sisi, _ = strconv.Atoi(r.URL.Query().Get("sisi"))
	jajarGenjang.Alas, _ = strconv.Atoi(r.URL.Query().Get("alas"))
	jajarGenjang.Tinggi, _ = strconv.Atoi(r.URL.Query().Get("tinggi"))
	jajarGenjang.Hitung = r.URL.Query().Get("hitung")
	if jajarGenjang.Hitung == "keliling" {
		go bangun_datar.KelilingJajarGenjang(hasilHitung, jajarGenjang.Alas, jajarGenjang.Sisi)
	} else if jajarGenjang.Hitung == "luas" {
		go bangun_datar.LuasJajarGenjang(hasilHitung, jajarGenjang.Alas, jajarGenjang.Tinggi)
	}
	a := <- hasilHitung
	kalimat := "Hasil Perhitungan: " + a
	utils.ResponseJSON(w, kalimat, http.StatusOK)
}

//CATEGORY

//Read
//GET Category
func GetCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    categories, err := categories.GetAll(ctx)
    if err != nil {
        fmt.Println(err)
    }
    utils.ResponseJSON(w, categories, http.StatusOK)
}

//GET Category with Filter
func GetFilter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    var idCategory = ps.ByName("id")
	var judul = r.URL.Query().Get("title")
	var minTahun = r.URL.Query().Get("minYear")
	var maxTahun = r.URL.Query().Get("maxYear")
	var minHalaman = r.URL.Query().Get("minPage")
	var maxHalaman = r.URL.Query().Get("maxPage")
	var urutkan = r.URL.Query().Get("sortByTitle")
	collect := []string{judul, minTahun, maxTahun, minHalaman, maxHalaman, urutkan}
	books, err := categories.GetCategoryFilter(ctx, idCategory, collect)
	if err != nil {
        fmt.Println(err)
    }
    utils.ResponseJSON(w, books, http.StatusOK)
}

// Create
// PostCategory
func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var category models.Category

    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    if err := categories.Insert(ctx, category); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)

}

//Update
// UpdateCategory
func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    var idCategory = ps.ByName("id")

    if err := categories.Update(ctx, category, idCategory); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteCategory
func DeleteCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idCategory = ps.ByName("id")

    if err := categories.Delete(ctx, idCategory); err != nil {
        kesalahan := map[string]string{
            "error": fmt.Sprintf("%v", err),
        }
        utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusOK)
}

//BOOKS

//Read
//GetBook
func GetBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var judul = r.URL.Query().Get("title")
	var minTahun = r.URL.Query().Get("minYear")
	var maxTahun = r.URL.Query().Get("maxYear")
	var minHalaman = r.URL.Query().Get("minPage")
	var maxHalaman = r.URL.Query().Get("maxPage")
	var urutkan = r.URL.Query().Get("sortByTitle")
	collect := []string{judul, minTahun, maxTahun, minHalaman, maxHalaman, urutkan}
	books, err := books.GetAll(ctx, collect)
	if err != nil {
	  fmt.Println(err)
	}
	utils.ResponseJSON(w, books, http.StatusOK)
}

// Create
// PostBook
func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	if err := books.Insert(ctx, book); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateBook
func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var book models.Book
  
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idBook = ps.ByName("id")
  
	if err := books.Update(ctx, book, idBook); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteBook
func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idBook = ps.ByName("id")
	if err := books.Delete(ctx, idBook); err != nil {
	  kesalahan := map[string]string{
		"error": fmt.Sprintf("%v", err),
	  }
	  utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

//MAIN FUNCTION
func main() {
	router := httprouter.New()
	router.GET("/bangun-datar/segitiga-sama-sisi", GetSegitiga)
	router.GET("/bangun-datar/persegi", GetPersegi)
	router.GET("/bangun-datar/persegi-panjang", GetPersegiPanjang)
	router.GET("/bangun-datar/lingkaran", GetLingkaran)
	router.GET("/bangun-datar/jajar-genjang", GetJajarGenjang)

	router.GET("/categories", GetCategory)
	router.GET("/categories/:id/books", GetFilter)
	router.POST("/categories", Auth(PostCategory))
	router.PUT("/categories/:id", Auth(UpdateCategory))
	router.DELETE("/categories/:id", Auth(DeleteCategory))

	router.GET("/books", GetBook)
	router.POST("/books", Auth(PostBook))
	router.PUT("/books/:id", Auth(UpdateBook))
	router.DELETE("/books/:id", Auth(DeleteBook))
	
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}