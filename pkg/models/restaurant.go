package models

import "database/sql"

type Restaurant struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Image        string       `json:"image"`
	WorkingHours WorkingHours `json:"workingHours"`
	Menu         []Product    `json:"menu"`
}

var _ Model = &Restaurant{}

func (r *Restaurant) Insert(db *sql.DB, params ...interface{}) error {
	query := `INSERT INTO Restaurant(id, name , image, open_at, close_at) 
			  VALUES (?,?,?,?,?) `
	_, err := db.Exec(query, r.Id, r.Name, r.Image,
		r.WorkingHours.Opening, r.WorkingHours.Closing)
	return err
}
func DropRestaurant(db *sql.DB) {
	query := "DELETE FROM Restaurant"
	db.Exec(query)
}
