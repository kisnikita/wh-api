package model

type Warehouse struct {
	ID        int        `jsonapi:"primary,warehouse"`
	Name      string     `jsonapi:"attr,name"`
	Available bool       `jsonapi:"attr,available"`
	Products  []*Product `jsonapi:"relation,product,omitempty"`
}

type Product struct {
	ID   int    `jsonapi:"primary,product"`
	Name string `jsonapi:"attr,name"`
	Code string `jsonapi:"attr,code"`
	Size int    `jsonapi:"attr,size"`
}

type WarehouseProduct struct {
	ID          int `jsonapi:"primary,warehouse_product" db:"id"`
	WarehouseID int `jsonapi:"attr,warehouse_id" db:"warehouse_id"`
	ProductID   int `jsonapi:"attr,product_id" db:"product_id"`
	Amount      int `jsonapi:"attr,amount" db:"amount"`
	Reserved    int `jsonapi:"attr,reserved" db:"reserved"`
}
