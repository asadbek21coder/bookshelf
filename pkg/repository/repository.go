package repository

import (
	"github.com/asadbek21coder/bookshelf"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user bookshelf.User) (bookshelf.User, error)
	GetUser(header bookshelf.Header) (bookshelf.User, error)
}

type Book interface {
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
