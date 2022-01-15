package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"taska/db"
	"taska/pkg/filesystem"
	"taska/pkg/models"
)

func ReadJSON(filePath string, c chan models.Restaurant) error {
	var restaurant models.Restaurant
	data, err := filesystem.ScanFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &restaurant)
	c <- restaurant
	return err
}
func DropTable(db *sql.DB)  {
	query :="DELETE FROM Restaurant"
	db.Exec(query)
}
func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	filePath := "../pkg/files/"
	var restaurants []models.Restaurant
	c := make(chan models.Restaurant, 7)
	for i := 0; i < 7; i++ {
		iterator := i+1
		go func() {
			err := ReadJSON(filePath+strconv.Itoa(iterator)+".json", c)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	for i := 0; i < 7; i++ {
		restaurants = append(restaurants, <-c)
	}

	DropTable(conn)

	var wg sync.WaitGroup
	for i := 0; i <7; i++ {
		iterator := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := restaurants[iterator].Insert(conn)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}
