package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kisnikita/wh-api/internal/model"
)

type Warehouse interface {
	GetAvailableProducts(warehouseId int) ([]*model.AvailableProductsResponse, error)
	ReserveProducts(reservations model.ReserveProductsRequest) ([]*model.UpdateProductsResponse, error)
	ReleaseProducts(releases model.ReleaseProductsRequest) ([]*model.UpdateProductsResponse, error)
}

type Repository struct {
	Warehouse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Warehouse: NewWarehousePostgres(db)}
}
