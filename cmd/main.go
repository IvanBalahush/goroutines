package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"sync"
	"taska/db"
	workerPool "taska/internal/worker_pool"
	"taska/pkg/filesystem"
	"taska/pkg/models"
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



	models.DropRestaurant(conn)
	models.DropProduct(conn)

	pool := workerPool.NewPool(4)
	count := pool.WorkerCount
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		//wg.Add(1)
		//go func(i int) {
		//	defer wg.Done()
		//	err := restaurants[i].Insert(conn)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}(i)
		wg.Add(1)
		go pool.Run(&wg, func(restaurant models.Restaurant) {
			restaurantID := models.GetID(conn, `SELECT id FROM Restaurant WHERE name = ?`, `INSERT INTO Restaurant(name) VALUE (?)`, restaurant.Name)
			_, err := conn.Exec(`INSERT INTO Restaurant VALUE (?,?,?,?,?)`, restaurant.Id, restaurant.Name, restaurant.Image, restaurant.WorkingHours.Opening, restaurant.WorkingHours.Closing)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(restaurantID)

		})
	}
	for _, restaurant := range restaurants {
		pool.Jobs <- restaurant
	}
	wg.Wait()

}
