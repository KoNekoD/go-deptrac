package controller

import "github.com/KoNekoD/go-deptrac/examples/simple_invalid_mvc/repository"

type Controller interface{}

type controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return &controller{
		repo: repo,
	}
}
