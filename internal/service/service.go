package service

import (
	"github.com/kisnikita/wh-api/internal/model"
	"github.com/kisnikita/wh-api/internal/repository"
)

type Warehouse interface {
	GetAvailableProducts(warehouseId int) ([]*model.AvailableProductsResponse, error)
	ReserveProducts(reservations model.ReserveProductsRequest) ([]*model.UpdateProductsResponse, error)
	ReleaseProducts(releases model.ReleaseProductsRequest) ([]*model.UpdateProductsResponse, error)
}

type Service struct {
	Warehouse
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Warehouse: NewWarehouseService(repo)}
}
