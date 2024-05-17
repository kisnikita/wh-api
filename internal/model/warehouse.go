package model

type Warehouse struct {
	ID        int    `jsonapi:"id"`
	Name      string `jsonapi:"name"`
	Available bool   `jsonapi:"available"`
}
