package service

import "github.com/kisnikita/wh-api/internal/repository"

type Warehouse interface {
}

type Service struct {
	Warehouse
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
