package model

type Reservation struct {
	ID     int    `jsonapi:"primary,reservation"`
	Code   string `jsonapi:"attr,code"`
	Amount int    `jsonapi:"attr,amount"`
}

type ReserveProductsRequest struct {
	ID           int            `jsonapi:"primary,message"`
	WarehouseID  int            `jsonapi:"attr,warehouse_id"`
	Reservations []*Reservation `jsonapi:"relation,reservations,omitempty"`
}

type ReleaseProductsRequest struct {
	ID          int            `jsonapi:"primary,message"`
	WarehouseID int            `jsonapi:"attr,warehouse_id"`
	Releases    []*Reservation `jsonapi:"relation,reservations,omitempty"`
}
