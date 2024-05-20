package model

type AvailableProductsResponse struct {
	ID     int `jsonapi:"primary,product"`
	Amount int `jsonapi:"attr,amount"`
}

type UpdateProductsResponse struct {
	ID       int `jsonapi:"primary,product"`
	Amount   int `jsonapi:"attr,amount"`
	Reserved int `jsonapi:"attr,reserved"`
}
