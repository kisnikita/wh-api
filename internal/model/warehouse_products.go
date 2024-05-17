package model

type WarehouseProduct struct {
	ID          int `jsonapi:"id"`
	WarehouseID int `jsonapi:"warehouse_id"`
	ProductID   int `jsonapi:"product_id"`
	Amount      int `jsonapi:"amount"`
	Reserved    int `jsonapi:"reserved"`
}
