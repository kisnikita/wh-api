package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kisnikita/wh-api/internal/model"
)

type WarehousePostgres struct {
	db *sqlx.DB
}

func NewWarehousePostgres(db *sqlx.DB) *WarehousePostgres {
	return &WarehousePostgres{db: db}
}

func (r *WarehousePostgres) GetAvailableProducts(warehouseId int) ([]*model.AvailableProductsResponse, error) {
	var products []model.AvailableProductsResponse
	query := fmt.Sprintf(`SELECT DISTINCT p.id, whp.amount 
								 FROM %s whp 
         						 JOIN %s wh ON whp.warehouse_id = $1
								 JOIN %s p ON whp.product_id = p.id`,
		warehouseProductsTable, warehousesTable, productsTable)
	err := r.db.Select(&products, query, warehouseId)

	if len(products) == 0 {
		return nil, errors.New("input id doesn't match to any warehouse")
	}

	response := make([]*model.AvailableProductsResponse, len(products))

	for i, product := range products {
		response[i] = &model.AvailableProductsResponse{
			ID:     product.ID,
			Amount: product.Amount,
		}
	}

	return response, err
}

func (r *WarehousePostgres) ReserveProducts(request model.ReserveProductsRequest) ([]*model.UpdateProductsResponse, error) {
	// Проверка, что введённые коды продуктов существуют в базе и их количества достаточно для резервации
	query := fmt.Sprintf(`SELECT * FROM %s whp
          						 WHERE whp.id IN (
								 	SELECT whp.id FROM %s 
         						 	WHERE whp.warehouse_id = $1 AND whp.product_id IN (
								 		SELECT ID FROM products p2 WHERE p2.code = $2
          						 	)
          						 )`,
		warehouseProductsTable, warehouseProductsTable)
	products := make([][]model.WarehouseProduct, len(request.Reservations))
	for i, reservation := range request.Reservations {
		err := r.db.Select(&products[i], query, request.WarehouseID, reservation.Code)
		if err != nil {
			return nil, err
		}
		if len(products[i]) == 0 {
			errTxt := fmt.Sprintf("input code = %s doesn't match to any product on warehouse with id = %d",
				reservation.Code, request.WarehouseID)
			return nil, errors.New(errTxt)
		}
		if products[i][0].Amount < reservation.Amount {
			errTxt := fmt.Sprintf("Not enough product with code = %s on warehouse with id = %d for reservation",
				reservation.Code, request.WarehouseID)
			return nil, errors.New(errTxt)
		}
	}

	// Внесения изменений в таблицу
	for _, res := range request.Reservations {
		query := fmt.Sprintf(`UPDATE %s whp SET amount = whp.amount - $1, reserved = whp.reserved + $1
          						 WHERE whp.amount >= $1 AND whp.id IN (
								 	SELECT whp.id FROM %s 
         						 	WHERE whp.warehouse_id = $2 AND whp.product_id IN (
								 		SELECT ID FROM products p2 WHERE p2.code = $3
          						 	)
          						 )`,
			warehouseProductsTable, warehouseProductsTable)
		_, err := r.db.Exec(query, res.Amount, request.WarehouseID, res.Code)
		if err != nil {
			return nil, err
		}
	}

	// Изменяем считанный для проверки массив
	for i, res := range request.Reservations {
		products[i][0].Amount -= res.Amount
		products[i][0].Reserved += res.Amount
	}

	// преобразуем данные в структуру ответа
	response := make([]*model.UpdateProductsResponse, len(products))
	for i, product := range products {
		response[i] = &model.UpdateProductsResponse{
			ID:       product[0].ID,
			Amount:   product[0].Amount,
			Reserved: product[0].Reserved,
		}
	}

	return response, nil

}

func (r *WarehousePostgres) ReleaseProducts(request model.ReleaseProductsRequest) ([]*model.UpdateProductsResponse, error) {
	// Проверка, что введённые коды продуктов существуют в базе и их количества достаточно для резервации
	query := fmt.Sprintf(`SELECT * FROM %s whp
          						 WHERE whp.id IN (
								 	SELECT whp.id FROM %s 
         						 	WHERE whp.warehouse_id = $1 AND whp.product_id IN (
								 		SELECT ID FROM products p2 WHERE p2.code = $2
          						 	)
          						 )`,
		warehouseProductsTable, warehouseProductsTable)
	products := make([][]model.WarehouseProduct, len(request.Releases))
	for i, release := range request.Releases {
		err := r.db.Select(&products[i], query, request.WarehouseID, release.Code)
		if err != nil {
			return nil, err
		}
		if len(products[i]) == 0 {
			errTxt := fmt.Sprintf("input code = %s doesn't match to any product on warehouse with id = %d",
				release.Code, request.WarehouseID)
			return nil, errors.New(errTxt)
		}
		if products[i][0].Reserved < release.Amount {
			errTxt := fmt.Sprintf("Not enough reservations on product with code = %s on warehouse with id = %d for release",
				release.Code, request.WarehouseID)
			return nil, errors.New(errTxt)
		}
	}

	// Внесения изменений в таблицу
	for _, res := range request.Releases {
		query := fmt.Sprintf(`UPDATE %s whp SET reserved = whp.reserved - $1
          						 WHERE whp.reserved >= $1 AND whp.id IN (
								 	SELECT whp.id FROM %s 
         						 	WHERE whp.warehouse_id = $2 AND whp.product_id IN (
								 		SELECT ID FROM products p2 WHERE p2.code = $3
          						 	)
          						 )`,
			warehouseProductsTable, warehouseProductsTable)
		_, err := r.db.Exec(query, res.Amount, request.WarehouseID, res.Code)
		if err != nil {
			return nil, err
		}
	}

	// Изменяем считанный для проверки массив
	for i, res := range request.Releases {
		products[i][0].Reserved -= res.Amount
	}

	// преобразуем данные в структуру ответа
	response := make([]*model.UpdateProductsResponse, len(products))
	for i, product := range products {
		response[i] = &model.UpdateProductsResponse{
			ID:       product[0].ID,
			Amount:   product[0].Amount,
			Reserved: product[0].Reserved,
		}
	}

	return response, nil
}
