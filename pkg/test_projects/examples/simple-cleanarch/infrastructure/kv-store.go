package infrastructure

import (
	"github.com/KoNekoD/go-deptrac/pkg/test_projects/examples/simple-cleanarch/domain"
)

type kvStore struct {
}

func (kv *kvStore) GetUser(id string) domain.User {
	return domain.User{
		Id:   id,
		Name: "John Doe",
	}
}

func (kv *kvStore) CreateUser(name string) domain.User {
	return domain.User{
		Id:   "001",
		Name: name,
	}
}
