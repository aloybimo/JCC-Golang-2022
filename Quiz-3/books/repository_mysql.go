package books

import (
  "quiz3/config"
  "quiz3/models"
  "database/sql"
  "net/url"
  "errors"
  "strings"
  "context"
  "fmt"
  "log"
  "time"
)

const (
  table          = "book"
  layoutDateTime = "2006-01-02 15:04:05"
)

// Read Books
func GetAll(ctx context.Context, param []string) ([]models.Book, error) {
	var books []models.Book
  	db, err := config.MySQL()
  	if err != nil {
    	log.Fatal("Cant connect to MySQL", err)
  	}
  	a := 0
  	for _, item := range param{
		if item != "" {
			a += 1
	  	}
  	}
  	queryText := "SELECT * FROM book WHERE "
	if a > 0{
		if param[0] != ""{
			queryText += "title = " + `"`+param[0]+`"` +" AND "
			a -= 1
		}
		if param[1] != "" {
			queryText += "release_year >= " + param[1] +" AND "
			a -= 1
		}
		if param[2] != "" {
			queryText += "release_year <= " + param[2] + " AND "
			a -= 1
		}
		if param[3] != "" {
			queryText += "total_page >= " + param[3] + " AND "
			a -= 1
		}
		if param[4] != "" {
			queryText += "total_page <= " + param[4] + " AND "
			a -= 1
		}
	} 

	s := strings.Split(queryText, " ")
	s = s[:len(s)-1]

	if param[5] != "" {
		if s[len(s)-1] == "AND" {
			queryText = strings.Join(s[:len(s)-1], " ")
			queryText += " ORDER BY title " + param[5] 
		} else {
			queryText += " ORDER BY title " + param[5]
		}
	} else {
		if s[len(s)-1] == "AND" || s[len(s)-1] == "WHERE" {
			queryText = strings.Join(s[:len(s)-1], " ") 
		} 
	}
  	
  	rowQuery, err := db.QueryContext(ctx, queryText)
  	if err != nil {
    	log.Fatal(err)
  	}
  	for rowQuery.Next() {
    	var book models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryID); err != nil {
      	return nil, err
    	}
		//  Change format string to datetime for created_at and updated_at
		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
      		log.Fatal(err)
    	}
    	book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
    	if err != nil {
      		log.Fatal(err)
    	}
    	books = append(books, book)
  	}
  	return books, nil
}

// Insert Books
func endApp() {
	message := recover()
  	fmt.Println("Terjadi Error", message)
}

func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	var thickness string
	if book.TotalPage <= 100 {
		thickness = "tipis"
	} else if book.TotalPage > 100 && book.TotalPage <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	defer endApp()
	_, err1 := url.ParseRequestURI(book.ImageURL)
	if err1 != nil {
		panic(err1)
	}

	defer endApp()
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		err2 := errors.New("Tahun Tidak Valid")
		panic(err2) 
	} 

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values('%v', '%v', '%v', %v, '%v', %v, '%v', NOW(), NOW(), %v)", table,
	  book.Title,
	  book.Description,
	  book.ImageURL,
	  book.ReleaseYear,
	  book.Price,
	  book.TotalPage,
	  thickness,
	  book.CategoryID)
	_, err = db.ExecContext(ctx, queryText)
  
	if err != nil {
	  return err
	}
	return nil
}

// Update Books
func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	var thickness string
	if book.TotalPage <= 100 {
		thickness = "tipis"
	} else if book.TotalPage > 100 && book.TotalPage <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	defer endApp()
	_, err1 := url.ParseRequestURI(book.ImageURL)
	if err1 != nil {
		panic(err1)
	}

	defer endApp()
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		err2 := errors.New("Tahun Tidak Valid")
		panic(err2) 
	} 

	queryText := fmt.Sprintf("UPDATE %v set title = '%v', description = '%v', image_url='%v', release_year=%v, price='%v', total_page=%v, thickness='%v', updated_at=NOW(), category_id=%v where id = %s", 
		table,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		thickness,
		book.CategoryID,
		id)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Delete Books
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