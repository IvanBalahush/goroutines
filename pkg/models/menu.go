package models

type Menu struct {
	Id int
	RestID int
	Product []Product
	Price int
}

//var _ Model = &Menu{}

//func (m *Menu) Insert(db *sql.DB, params ...interface{}) error {
//	query := `INSERT INTO Menu( id, rest_id, product_id, price)
//			  VALUES (?,?,?,?)`
//	_, err := db.Exec(query, m.ID)
//	return err
//}
