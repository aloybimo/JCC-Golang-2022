package mahasiswa


import (
  "tugas15/config"
  "tugas15/models"
  "context"
  "database/sql"
  "errors"
  "fmt"
  "log"
  "time"
)
 
const (
  table          = "mahasiswa"
  layoutDateTime = "2006-01-02 15:04:05"
)
 
// GetAll Mahasiswa
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
  var dataMahasiswa []models.Mahasiswa
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Cant connect to MySQL", err)
  }
  queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
  rowQuery, err := db.QueryContext(ctx, queryText)
  if err != nil {
    log.Fatal(err)
  }
  for rowQuery.Next() {
    var mahasiswa models.Mahasiswa
    var createdAt, updatedAt string
    if err = rowQuery.Scan(&mahasiswa.ID,
      &mahasiswa.Nama,
      &createdAt,
      &updatedAt); err != nil {
      return nil, err
    }
    //  Change format string to datetime for created_at and updated_at
    mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
    if err != nil {
      log.Fatal(err)
    }
    mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
    if err != nil {
      log.Fatal(err)
    }
    dataMahasiswa = append(dataMahasiswa, mahasiswa)
  }
  return dataMahasiswa, nil
}
 
// Insert Mahasiswa
func Insert(ctx context.Context, mahasiswa models.Mahasiswa) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  queryText := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) values('%v', NOW(), NOW())", table,
    mahasiswa.Nama)
  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }
  return nil
}
 
// Update Mahasiswa
func Update(ctx context.Context, mahasiswa models.Mahasiswa, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  queryText := fmt.Sprintf("UPDATE %v set nama ='%s', updated_at = NOW() where id = %s",
    table,
    mahasiswa.Nama,
    id,
  )
  fmt.Println(queryText)
  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }
  return nil
}
 
// Delete Mahasiswa
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