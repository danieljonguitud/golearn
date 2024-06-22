package cmdmanager

import "fmt"

type CMDManager struct {}

func (cmdm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdm CMDManager) WriteData(data interface{}) (error) {}

func New() CMDManager {
	return CMDManager{}
}
