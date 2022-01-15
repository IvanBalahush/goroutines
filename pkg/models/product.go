package models

type Product struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Price       float64      `json:"price"`
	Image       string       `json:"image"`
	Type        string       `json:"type"`
	Ingredients []string     `json:"ingredients,omitempty"`
}