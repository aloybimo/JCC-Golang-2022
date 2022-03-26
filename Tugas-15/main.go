package main 

import(
	"tugas15/nilai"
	"tugas15/mahasiswa"
	"tugas15/mata_kuliah"
	"tugas15/utils"
	"tugas15/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
) 

// Read
// GetNilai
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilaiNilai, err := nilai.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, nilaiNilai, http.StatusOK)
}

// Create
// PostNilai
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var satuNilai models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&satuNilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	if err := nilai.Insert(ctx, satuNilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateNilai
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var satuNilai models.NilaiMahasiswa
  
	if err := json.NewDecoder(r.Body).Decode(&satuNilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idNilai = ps.ByName("id")
  
	if err := nilai.Update(ctx, satuNilai, idNilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteNilai
func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := nilai.Delete(ctx, idNilai); err != nil {
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

// Read
// GetMataKuliah
func GetMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    dataMataKuliah, err := mata_kuliah.GetAll(ctx)
    if err != nil {
        fmt.Println(err)
    }
    utils.ResponseJSON(w, dataMataKuliah, http.StatusOK)
}

// Create
// PostMataKuliah
func PostMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var mataKuliah models.MataKuliah

    if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    if err := mata_kuliah.Insert(ctx, mataKuliah); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)

}

//Update
// UpdateMataKuliah
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var mataKuliah models.MataKuliah

    if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    var idMataKuliah = ps.ByName("id")

    if err := mata_kuliah.Update(ctx, mataKuliah, idMataKuliah); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMataKuliah
func DeleteMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idMataKuliah = ps.ByName("id")

    if err := mata_kuliah.Delete(ctx, idMataKuliah); err != nil {
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

// Read
// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    dataMahasiswa, err := mahasiswa.GetAll(ctx)
    if err != nil {
        fmt.Println(err)
    }
    utils.ResponseJSON(w, dataMahasiswa, http.StatusOK)
}

// Create
// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var satuMahasiswa models.Mahasiswa

    if err := json.NewDecoder(r.Body).Decode(&satuMahasiswa); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    if err := mahasiswa.Insert(ctx, satuMahasiswa); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)

}

// Update
// UpdateMahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var satuMahasiswa models.Mahasiswa

    if err := json.NewDecoder(r.Body).Decode(&satuMahasiswa); err != nil {
        utils.ResponseJSON(w, err, http.StatusBadRequest)
        return
    }

    var idMahasiswa = ps.ByName("id")

    if err := mahasiswa.Update(ctx, satuMahasiswa, idMahasiswa); err != nil {
        utils.ResponseJSON(w, err, http.StatusInternalServerError)
        return
    }

    res := map[string]string{
        "status": "Succesfully",
    }

    utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMahasiswa
func DeleteMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idMahasiswa = ps.ByName("id")

    if err := mahasiswa.Delete(ctx, idMahasiswa); err != nil {
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

func main() {
	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa/create", PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)

	router.GET("/matakuliah", GetMataKuliah)
	router.POST("/matakuliah/create", PostMataKuliah)
	router.PUT("/matakuliah/:id/update", UpdateMataKuliah)
	router.DELETE("/matakuliah/:id/delete", DeleteMataKuliah)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}