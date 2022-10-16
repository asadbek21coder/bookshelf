package service

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/asadbek21coder/bookshelf"
	"github.com/asadbek21coder/bookshelf/pkg/repository"
)

type Authorization interface {
	CreateUser(user bookshelf.User) (bookshelf.User, error)
	GetUser(header bookshelf.Header) (bookshelf.User, error)
}

type Book interface {
	CreateBook(isbn bookshelf.Isbn, header bookshelf.Header) (bookshelf.BookDB, error)
	GetUserSecret(header bookshelf.Header) (string, error)
	GetAllBooks(header bookshelf.Header) ([]bookshelf.FullBook, error)
}

type Service struct {
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book),
	}
}

func generateSignMD5Hash(payload string) string {

	hash := md5.Sum([]byte(payload))
	return hex.EncodeToString(hash[:])
}
