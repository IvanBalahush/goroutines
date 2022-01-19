package db

type Ingredient struct {
	ID   int
	Name string
}

type Menu struct {
	ID     int
	RestID int
}
type MenuProducts struct {
	ProductID int
	MenuID    int
	Price     float64
}
type ProductIngredients struct {
	ProductID    int
	IngredientID int
}
type Product struct {
	ID    int
	Name  string
	Image string
	Type  string
}
type Restaurant struct {
	ID    int
	Name  string
	Image string
	Open  string
	Close string
}
