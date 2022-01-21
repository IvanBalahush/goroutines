package main

import (
	"fmt"
	"log"
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
	err = models.DeleteTables(conn)
	if err != nil {
		log.Fatal(err)
	}

	pool := workerPool.NewWorkerPool(4)
	var wg sync.WaitGroup
	wg.Add(pool.Count)

	for i := 0; i < pool.Count; i++ {
		go pool.Run(&wg, func(restaurant models.Restaurant) {
			fmt.Println("Start")
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
				if err != nil {
					fmt.Println(err)
				}
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
	filePath := "../assets/files/"
	restaurants := [7]models.Restaurant{}

	for i := 1; i <= 7; i++ {
		rest, err := filesystem.ParseFiles(filePath + strconv.Itoa(i) + ".json")
		if err != nil {
			log.Fatal(err)
		}
		restaurants[i-1] = rest
		pool.Sender <- rest
	}

	pool.Stop()
	wg.Wait()
}
