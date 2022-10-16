package repository

import (
	"fmt"

	"github.com/asadbek21coder/bookshelf"
	"github.com/jmoiron/sqlx"
)

const (
	shelfTable = "shelf"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) CreateBook(isbn string, header bookshelf.Header) (bookshelf.BookDB, error) {
	var resp bookshelf.BookDB
	query := fmt.Sprintf(`INSERT INTO %s (isbn,status,user_key) values ($1,$2, $3) RETURNING *`, shelfTable)
	row, err := r.db.Query(query, isbn, 0, *header.Key)
	if err != nil {
		return bookshelf.BookDB{}, err
	}
	for row.Next() {
		err := row.Scan(
			&resp.Id,
			&resp.Isbn,
			&resp.Status,
			&resp.UserKey,
		)

		if err != nil {
			return bookshelf.BookDB{}, err
		}
	}

	return resp, nil
}

// func (r *BookPostgres) GetBook(header bookshelf.Header) (bookshelf.Book, error) {
// 	// isbn := "9781118464465"
// 	// url := "https://openlibrary.org/api/books?bibkeys=ISBN:" + isbn + "&jscmd=details&format=json"
// 	// response, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("error")
// 	// 	return bookshelf.Book{}, err
// 	// }
// 	// responseData, err := ioutil.ReadAll(response.Body)
// 	// if err != nil {
// 	// 	fmt.Println("error")
// 	// 	return bookshelf.Book{}, err
// 	// }
// 	// if err != nil {
// 	// 	fmt.Println("error")
// 	// 	return bookshelf.Book{}, err
// 	// }
// 	// fmt.Println(responseData)
// 	return bookshelf.Book{}, nil
// }

func (r *BookPostgres) GetUserSecret(header bookshelf.Header) (string, error) {
	var userSecret string

	query := fmt.Sprintf(`SELECT secret FROM %s WHERE key=$1`, usersTable)
	err := r.db.Get(&userSecret, query, *header.Key)
	return userSecret, err
}

func (r *BookPostgres) GetAllBooks(header bookshelf.Header) (response []bookshelf.FullBook, err error) {

	var resp []bookshelf.BookDB

	var data bookshelf.FullBook
	fmt.Printf("%T", data)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_key=$1`, shelfTable)
	err = r.db.Select(&resp, query, *header.Key)
	for i, v := range resp {
		fmt.Println(i, v)
		data.Book.Id = v.Id
		data.Book.Isbn = v.Isbn
		data.Status = v.Status
		response = append(response, data)
		// fmt.Println(v)
	}
	fmt.Println(resp)
	fmt.Println(data)
	if err != nil {
		return response, err
	}
	return response, nil

}
