package model

type Product struct {
	ID   int    `jsonapi:"id"`
	Name string `jsonapi:"name"`
	Code string `jsonapi:"code"`
	Size int    `jsonapi:"size"`
}
