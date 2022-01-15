package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"taska/db"
	"taska/pkg/filesystem"
	"taska/pkg/models"
)

func Draft()  {
	for i := 1; i <=7; i++ {
		filePath := "../pkg/files/" + strconv.Itoa(i) + ".json"
		data, err := filesystem.ScanFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		var menu models.Menu
		err = json.Unmarshal(data, &menu)
		if err != nil {
			log.Fatal(err)
		}
		//for _, product := range menu.Menu {
		//	fmt.Println("ID: ", product.ID)
		//	fmt.Println("Name: ", product.Name)
		//	fmt.Println("Price: ", product.Price)
		//	fmt.Println("Image: ", product.Image)
		//	fmt.Println("Type: ", product.Type)
		//	fmt.Println("ingredients: ", product.Ingredients)
		//}
	}
}
func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	filePath := "../pkg/files/1.json"
	data, err := filesystem.ScanFile(filePath)
	//var menu models.Menu
	//err = json.Unmarshal(data, &menu)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("ID: ", menu.Menu[0].ID)
	//fmt.Println("Name: ", menu.Menu[0].Name)
	//fmt.Println("Price: ", menu.Menu[0].Price)
	//fmt.Println("Image: ", menu.Menu[0].Image)
	//fmt.Println("Type: ", menu.Menu[0].Type)
	//fmt.Println("ingredients: ", menu.Menu[0].Ingredients)

	var restaurant models.Restaurant
	err = json.Unmarshal(data, &restaurant)
	fmt.Println(restaurant)
}
