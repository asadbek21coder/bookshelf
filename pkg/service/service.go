package service

import "github.com/asadbek21coder/bookshelf/pkg/repository"

type Authorization interface {
}

type Book interface {
}

type Service struct {
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
