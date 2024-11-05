package service

import "github.com/KoNekoD/go-deptrac/examples/simple_mvc/repository"

type Service interface{}

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}
