package models

import "database/sql"

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}

var _ Model = &Product{}

func (p Product) Insert(db *sql.DB, params ...interface{}) error {
	query := `INSERT INTO Product(id, name)
			  VALUES (?,?)`
	_, err := db.Exec(query, p.Id, p.Name)
	return err
}

func DropProduct(db *sql.DB) {
	query := "DELETE FROM Product"
	db.Exec(query)
}