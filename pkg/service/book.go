package service

import (
	"errors"

	"github.com/asadbek21coder/bookshelf"
	"github.com/asadbek21coder/bookshelf/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(isbn bookshelf.Isbn, header bookshelf.Header) (bookshelf.BookDB, error) {

	secret, err := s.repo.GetUserSecret(header)
	payload := "POST" + bookshelf.URL + "/books{\"isbn\":\"" + isbn.Isbn + "\"}" + secret
	if err != nil {
		return bookshelf.BookDB{}, err
	}
	hash := generateSignMD5Hash(payload)
	if hash != *header.Sign {
		return bookshelf.BookDB{}, errors.New("invalid sign")
	}

	return s.repo.CreateBook(isbn.Isbn, header)
}

func (s *BookService) GetUserSecret(header bookshelf.Header) (string, error) {
	return "", nil
}

func (s *BookService) GetAllBooks(header bookshelf.Header) ([]bookshelf.FullBook, error) {
	secret, err := s.repo.GetUserSecret(header)
	payload := "GET" + bookshelf.URL + "/books" + secret
	if err != nil {
		return []bookshelf.FullBook{}, err
	}
	hash := generateSignMD5Hash(payload)
	if hash != *header.Sign {
		return []bookshelf.FullBook{}, errors.New("invalid sign")
	}

	data, err := s.repo.GetAllBooks(header)
	if err != nil {
		return []bookshelf.FullBook{}, err
	}
	// fmt.Println(data)
	return data, nil
}
