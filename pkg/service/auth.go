package service

import (
	"errors"

	"github.com/asadbek21coder/bookshelf"
	"github.com/asadbek21coder/bookshelf/pkg/repository"
)

// const key = "MyUserKey"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user bookshelf.User) (bookshelf.User, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(header bookshelf.Header) (bookshelf.User, error) {
	data, err := s.repo.GetUser(header)
	payload := "GET" + bookshelf.URL + "/myself" + data.Secret
	if err != nil {
		return bookshelf.User{}, err
	}
	hash := generateSignMD5Hash(payload)
	if hash != *header.Sign {
		return bookshelf.User{}, errors.New("invalid sign")
	}
	return data, nil
}
