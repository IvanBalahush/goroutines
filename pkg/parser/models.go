package parser

import (
	"database/sql"
	"fmt"
)

type Restaurant struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Image        string `json:"image"`
	WorkingHours struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"workingHours"`
	Menu []Menu `json:"menu"`
}
type Menu struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}

func (r *Restaurant) Insert(db *sql.DB) (int64, error) {
	query := `INSERT INTO Restaurant (name, image, open, close)
			  VALUES (?,?,?,?)`
	res, err := db.Exec(query, r.Name, r.Image, r.WorkingHours.Opening, r.WorkingHours.Closing)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (p Menu) Insert(db *sql.DB) (int64, error) {
	query := `INSERT INTO products( name, price,image, type) 
			  VALUES (?,?,?,?)`
	res, err := db.Exec(query, p.Name, p.Price, p.Image, p.Type)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func DeleteTables(db *sql.DB) error {
	query := `DELETE FROM Restaurant`
	_, err := db.Exec(query)
	q2 := `DELETE FROM products`
	_, err = db.Exec(q2)
	q3 := `DELETE FROM ingredients`
	_, err = db.Exec(q3)
	q4 := `DELETE FROM rest_products`
	_, err = db.Exec(q4)
	q5 := `DELETE FROM product_ingredients`
	_, err = db.Exec(q5)
	return err
}

func (w Restaurant) Do(data interface{}, i int) {
	fmt.Printf("Goroutine number %d, is running. Input %s \n", i, data)
}
func (w Restaurant) Stop() {

}
