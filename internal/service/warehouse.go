package service

import (
	"github.com/kisnikita/wh-api/internal/model"
	"github.com/kisnikita/wh-api/internal/repository"
)

type WarehouseService struct {
	repo repository.Warehouse
}

func NewWarehouseService(repo repository.Warehouse) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) GetAvailableProducts(warehouseId int) ([]*model.AvailableProductsResponse, error) {
	return s.repo.GetAvailableProducts(warehouseId)
}

func (s *WarehouseService) ReserveProducts(reservations model.ReserveProductsRequest) ([]*model.UpdateProductsResponse, error) {
	return s.repo.ReserveProducts(reservations)
}

func (s *WarehouseService) ReleaseProducts(releases model.ReleaseProductsRequest) ([]*model.UpdateProductsResponse, error) {
	return s.repo.ReleaseProducts(releases)
}
