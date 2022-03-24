package main 

import(
	"tugas14/nilai"
	"tugas14/utils"
	"tugas14/models"
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

func main() {
	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}