package main

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"taska/db"
	"taska/pkg/filesystem"
	models "taska/pkg/parser"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	filePath := "../assets/files/"

	var restaurants []models.Restaurant
	ch := make(chan models.Restaurant, 7)

	for i := 1; i <= 7; i++ {
		go func(i int) {
			err := filesystem.ReadJSON(filepath.Join(filePath, strconv.Itoa(i)+".json"), ch)
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	for i := 0; i < 7; i++ {
		restaurants = append(restaurants, <-ch)
	}
	err = models.DeleteTables(conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, restaurant := range restaurants {
		resID, err := restaurant.Insert(conn)
		if err != nil {
			log.Fatal(err)
		}

		for _, product := range restaurant.Menu {
			prodID, err := product.Insert(conn)
			if err != nil {
				if strings.HasPrefix(err.Error(), "Error 1062") {
					continue
				} else {
					log.Fatal(err)
				}
			}
			restProductQuery := `INSERT INTO rest_products(rest_id, product_id, price) 
					  VALUES (?,?,?)`
			_, err = conn.Exec(restProductQuery, resID, prodID, product.Price)
			for _, ingredient := range product.Ingredients {
				ingredientsQuery := `INSERT INTO ingredients(name)
				VALUES (?)`
				ingredientRes, err := conn.Exec(ingredientsQuery, ingredient)
				if err != nil {
					if strings.HasPrefix(err.Error(), "Error 1062") {
						continue
					} else {
						log.Fatal(err)
					}
				}
				ingID, err := ingredientRes.LastInsertId()
				if err != nil {
					log.Fatal(err)
				}
				productIngredientsQuery := `INSERT INTO product_ingredients (product_id, ingredient_id) 
											VALUES (?,?)`
				_, err = conn.Exec(productIngredientsQuery, prodID, ingID)
				if err != nil {
					if strings.HasPrefix(err.Error(), "Error 1062") {
						continue
					} else {
						log.Fatal(err)
					}
				}
			}
		}
	}
}
