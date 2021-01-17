package service

import (
	"github.com/mxmntv/todo_app/pkg/repository"
)

/*Auther - */
type Auther interface {
}

/*Lister - */
type Lister interface {
}

/*Itemer - */
type Itemer interface {
}

/*Service - */
type Service struct {
	Auther
	Lister
	Itemer
}

/*NewService - */
func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
