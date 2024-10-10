package controller

import "github.com/KoNekoD/go-deptrac/examples/simple_mvc/service"

type Controller interface{}

type controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return &controller{
		service: service,
	}
}
