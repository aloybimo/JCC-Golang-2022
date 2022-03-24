package nilai

import (
  "tugas14/config"
  "tugas14/models"
  "database/sql"
  "errors"
  "context"
  "fmt"
  "log"
  "time"
)

const (
  table          = "nilai"
  layoutDateTime = "2006-01-02 15:04:05"
)

// Read GetAll
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {
  var nilaiNilai []models.NilaiMahasiswa
  db, err := config.MySQL()

  if err != nil {
    log.Fatal("Cant connect to MySQL", err)
  }

  queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
  rowQuery, err := db.QueryContext(ctx, queryText)

  if err != nil {
    log.Fatal(err)
  }

  for rowQuery.Next() {
    var nilai models.NilaiMahasiswa
    var createdAt, updatedAt string
    if err = rowQuery.Scan(&nilai.ID,
      &nilai.Nama,
      &nilai.MataKuliah,
	  &nilai.IndeksNilai,
	  &nilai.Nilai,
      &createdAt,
      &updatedAt); err != nil {
      return nil, err
    }

    //  Change format string to datetime for created_at and updated_at
    nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

    if err != nil {
      log.Fatal(err)
    }

    nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
    if err != nil {
      log.Fatal(err)
    }

    nilaiNilai = append(nilaiNilai, nilai)
  }
  return nilaiNilai, nil
}

// Insert Nilai
func Insert(ctx context.Context, nilai models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	var indeks string
	if nilai.Nilai >= 80 && nilai.Nilai <= 100 {
		indeks = "A"
	} else if nilai.Nilai >= 70 && nilai.Nilai < 80 {
		indeks = "B"
	} else if nilai.Nilai >= 60 && nilai.Nilai < 70 {
		indeks = "C"
	} else if nilai.Nilai >= 50 && nilai.Nilai < 60 {
		indeks = "D"
	} else if nilai.Nilai < 50 {
		indeks = "E"
	} else {
		indeks = "-"
	}
	queryText := fmt.Sprintf("INSERT INTO %v (nama, mataKuliah, indeksNilai, nilai, created_at, updated_at) values('%v','%v', '%v', %v, NOW(), NOW())", table,
	  nilai.Nama,
	  nilai.MataKuliah,
	  indeks,
	  nilai.Nilai)
	_, err = db.ExecContext(ctx, queryText)
  
	if err != nil {
	  return err
	}
	return nil
}

// Update Nilai
func Update(ctx context.Context, nilai models.NilaiMahasiswa, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }

  var indeks string
	if nilai.Nilai >= 80 && nilai.Nilai <= 100 {
		indeks = "A"
	} else if nilai.Nilai >= 70 && nilai.Nilai < 80 {
		indeks = "B"
	} else if nilai.Nilai >= 60 && nilai.Nilai < 70 {
		indeks = "C"
	} else if nilai.Nilai >= 50 && nilai.Nilai < 60 {
		indeks = "D"
	} else if nilai.Nilai < 50 {
		indeks = "E"
	} else {
		indeks = "-"
	}

  queryText := fmt.Sprintf("UPDATE %v set nama ='%s', mataKuliah = '%d', indeksNilai = '%d', nilai = %d, updated_at = NOW() where id = %s",
    table,
    nilai.Nama,
    nilai.MataKuliah,
    indeks,
    nilai.Nilai,
    id,
  )

  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }

  return nil
}

// Delete Nilai
func Delete(ctx context.Context, id string) error {
  db, err := config.MySQL()
  if err != nil {
      log.Fatal("Can't connect to MySQL", err)
  }

  queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

  s, err := db.ExecContext(ctx, queryText)

  if err != nil && err != sql.ErrNoRows {
      return err
  }

  check, err := s.RowsAffected()
  fmt.Println(check)
  if check == 0 {
      return errors.New("id tidak ada")
  }

  if err != nil {
      fmt.Println(err.Error())
  }

  return nil
}