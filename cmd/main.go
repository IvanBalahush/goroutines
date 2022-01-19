package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"taska/db"
	workerPool "taska/internal/worker_pool"
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

	//
	pool := workerPool.NewPool(4)
	var wg sync.WaitGroup
	wg.Add(pool.WorkerCount)
	for i := 0; i < pool.WorkerCount; i++ {
		fmt.Println("Start")
		go pool.Run(&wg, func(restaurant models.Restaurant) {
			resID, err := restaurant.Insert(conn)
			if err != nil {
				log.Fatal(err)
			}
			selectExternalIDRestaurantQuery := `SELECT external_id 
												FROM Restaurant
												WHERE id = ?`
			RestExternalID := conn.QueryRow(selectExternalIDRestaurantQuery, resID)
			for _, product := range restaurant.Menu {
				prodID, err := product.Insert(conn)
				if err != nil {
					if strings.HasPrefix(err.Error(), "Error 1062") {
						continue
					} else {
						log.Fatal(err)
					}
				}
				//
				selectExternalIDProductQuery := `SELECT external_id
												 FROM products
												 WHERE id = ?`
				ProdExternalID := conn.QueryRow(selectExternalIDProductQuery, prodID)
				//
				restProductQuery := `INSERT INTO rest_products(rest_id, product_id, price, external_rest_id, external_product_id) 
					  VALUES (?,?,?)`
				_, err = conn.Exec(restProductQuery, resID, prodID, product.Price, RestExternalID, ProdExternalID)
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
		})

	}
	//for _, restaurant := range restaurants {
	//	resID, err := restaurant.Insert(conn)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	for _, product := range restaurant.Menu {
	//		prodID, err := product.Insert(conn)
	//		if err != nil {
	//			if strings.HasPrefix(err.Error(), "Error 1062") {
	//				continue
	//			} else {
	//				log.Fatal(err)
	//			}
	//		}
	//		restProductQuery := `INSERT INTO rest_products(rest_id, product_id, price)
	//				  VALUES (?,?,?)`
	//		_, err = conn.Exec(restProductQuery, resID, prodID, product.Price)
	//		for _, ingredient := range product.Ingredients {
	//			ingredientsQuery := `INSERT INTO ingredients(name)
	//			VALUES (?)`
	//			ingredientRes, err := conn.Exec(ingredientsQuery, ingredient)
	//			if err != nil {
	//				if strings.HasPrefix(err.Error(), "Error 1062") {
	//					continue
	//				} else {
	//					log.Fatal(err)
	//				}
	//			}
	//			ingID, err := ingredientRes.LastInsertId()
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			productIngredientsQuery := `INSERT INTO product_ingredients (product_id, ingredient_id)
	//										VALUES (?,?)`
	//			_, err = conn.Exec(productIngredientsQuery, prodID, ingID)
	//			if err != nil {
	//				if strings.HasPrefix(err.Error(), "Error 1062") {
	//					continue
	//				} else {
	//					log.Fatal(err)
	//				}
	//			}
	//		}
	//	}
	//}
}
