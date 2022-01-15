package filesystem

import (
	"bufio"
	"encoding/json"
	"os"
	"taska/pkg/models"
)

func Parse(inf []byte, s models.Menu) (models.Menu, error) {
	err := json.Unmarshal(inf, &s)

	return s, err
}

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