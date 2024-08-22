package controller

import (
	"github.com/KoNekoD/go-deptrac/pkg/test_projects/examples/simple-invalid-mvc/repository"
)

type Controller interface{}

type controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return &controller{
		repo: repo,
	}
}
