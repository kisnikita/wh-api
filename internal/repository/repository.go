package repository

import "github.com/jmoiron/sqlx"

type Warehouse interface {
}

type Repository struct {
	Warehouse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
