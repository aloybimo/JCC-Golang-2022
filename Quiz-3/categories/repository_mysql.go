package categories

import (
  "quiz3/config"
  "quiz3/models"
  "context"
  "database/sql"
  "errors"
  "strings"
  "fmt"
  "log"
  "time"
)
 
const (
  table          = "category"
  layoutDateTime = "2006-01-02 15:04:05"
)
 
// GetAll Categories
func GetAll(ctx context.Context) ([]models.Category, error) {
  var categories []models.Category
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
    var category models.Category
    var createdAt, updatedAt string
    if err = rowQuery.Scan(&category.ID,
      &category.Name,
      &createdAt,
      &updatedAt); err != nil {
      return nil, err
    }
    //  Change format string to datetime for created_at and updated_at
    category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
    if err != nil {
      log.Fatal(err)
    }
    category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
    if err != nil {
      log.Fatal(err)
    }
    categories = append(categories, category)
  }
  return categories, nil
}

// Get Categories with Filter
func GetCategoryFilter(ctx context.Context, id string, param []string) ([]models.Book, error) {
  var books []models.Book
  db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
  }
  a := 0
  for _, item := range param{
	  if item != "" {
			a += 1
	  }
  }
  queryText := fmt.Sprintf("SELECT * FROM book where category_id = %s AND", id)
	if a > 0{
		if param[0] != ""{
			queryText += " title = " + `"`+param[0]+`"` +" AND "
			a -= 1
		}
		if param[1] != "" {
			queryText += " release_year >= " + param[1] +" AND "
			a -= 1
		}
		if param[2] != "" {
			queryText += " release_year <= " + param[2] + " AND "
			a -= 1
		}
		if param[3] != "" {
			queryText += " total_page >= " + param[3] + " AND "
			a -= 1
		}
		if param[4] != "" {
			queryText += " total_page <= " + param[4] + " AND "
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
		} else {
      queryText = strings.Join(s[:len(s)], " ")
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

// Insert Categories
func Insert(ctx context.Context, category models.Category) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  queryText := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())", table,
    category.Name)
  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }
  return nil
}
 
// Update Categories
func Update(ctx context.Context, category models.Category, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  queryText := fmt.Sprintf("UPDATE %v set name ='%s', updated_at = NOW() where id = %s",
    table,
    category.Name,
    id,
  )
  fmt.Println(queryText)
  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }
  return nil
}
 
// Delete Categories
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