package filesystem

import (
	"bufio"
	"encoding/json"
	"os"
	models "taska/pkg/parser"
)

func ScanFile(filePath string) (data []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text()...)
	}
	return data, err
}

func ReadJSON(filePath string, c chan models.Restaurant) error {
	var restaurant models.Restaurant
	data, err := ScanFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &restaurant)
	c <- restaurant
	return err
}
