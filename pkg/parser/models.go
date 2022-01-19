package models

import "database/sql"

type Restaurant struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Image        string `json:"image"`
	WorkingHours struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"workingHours"`
	Menu []Product `json:"menu"`
}
type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}

func (r *Restaurant) Insert(db *sql.DB) error {
	query := `INSERT INTO Restaurant (id, name, image, open, close)
			  VALUES (?,?,?,?,?)`
	_, err := db.Exec(query, r.Id, r.Name, r.Image, r.WorkingHours.Opening, r.WorkingHours.Closing)

	return err
}
func (p Product) Insert(db *sql.DB) error {
	query := `INSERT INTO products(id, name, price, image, type) 
			  VALUES (?,?,?,?,?)`
	_, err := db.Exec(query, p.Id, p.Name, p.Price, p.Image, p.Type)
	return err
}
func DeleteRestaurant(db *sql.DB) error {
	query := `DELETE FROM Restaurant`
	_, err := db.Exec(query)
	return err
}

func DeleteProduct(db *sql.DB) error {
	query := `DELETE FROM products`
	_, err := db.Exec(query)
	return err
}
