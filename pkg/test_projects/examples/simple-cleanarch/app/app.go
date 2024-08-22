package app

import (
	"github.com/KoNekoD/go-deptrac/pkg/test_projects/examples/simple-cleanarch/domain"
)

type App interface {
	RegisterUser()
}

type app struct {
	repo domain.UserRepo
}

func New(repo domain.UserRepo) App {
	return &app{
		repo: repo,
	}
}

func (a *app) RegisterUser() {}
