package repository

import (
	"fmt"

	"github.com/asadbek21coder/bookshelf"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user bookshelf.User) (bookshelf.User, error) {
	var resp bookshelf.User
	// var id int
	query := fmt.Sprintf(`INSERT INTO %s (name,key,secret ) values ($1,$2, $3) RETURNING *`, usersTable)
	row, err := r.db.Query(query, user.Name, user.Key, user.Secret)
	if err != nil {
		return bookshelf.User{}, err
	}
	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
			&resp.Key,
			&resp.Secret,
		)
		if err != nil {
			return bookshelf.User{}, err
		}
	}

	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }
	return resp, nil
}

func (r *AuthPostgres) GetUser(header bookshelf.Header) (bookshelf.User, error) {
	var resp bookshelf.User
	query := fmt.Sprintf(`SELECT * from %s WHERE key=$1`, usersTable)
	row, err := r.db.Query(query, *header.Key)
	if err != nil {
		return bookshelf.User{}, err
	}
	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
			&resp.Key,
			&resp.Secret,
		)
		if err != nil {
			return bookshelf.User{}, err
		}
	}

	return resp, nil
}
