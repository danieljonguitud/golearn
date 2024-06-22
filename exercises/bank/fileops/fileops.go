package fileops 

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func WriteFloatOnFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("File not found")
	}
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}
