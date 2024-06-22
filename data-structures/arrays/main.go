package main

import "fmt"

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.9, 9.99, 45.9, 20.0}
// 	fmt.Println(productNames)
// 
// 	featuredPrices := prices[1:] // [9.99, 45.9, 20.0]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1] // [199.99]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// 
// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

type Product struct {
	id string
	title string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	// fmt.Println(prices[0:1])
	prices[1] = 9.99

	// 1
	hobbies := [3]string{"programming", "boxing", "videogames"}
	//fmt.Println(hobbies)

	// 2
	//fmt.Println(hobbies[0])
	//twoHobbies := hobbies[1:]
	//fmt.Println(twoHobbies)

	// 3
	firstTwoHobbies := hobbies[:2]
	fmt.Println(firstTwoHobbies)

	//sameHobbies := hobbies[0:2]
	//fmt.Println(sameHobbies)

	// 4
	reSliceHobbies := firstTwoHobbies[1:3]
	fmt.Println(reSliceHobbies)

	// 5
	var dynamicArr []string
	dynamicArr = []string{"learn go", "practice go"}

	// 6
	dynamicArr[1] = "more go"
	dynamicArr = append(dynamicArr, "third goal")
	fmt.Println(dynamicArr)

	// 7

	products := []Product{
		{
			"first-product",
			"A first product",
			12.99,
		},
		{
			"second-product",
			"A second product",
			12.99,
		},
		{
			"second-product",
			"A second product",
			12.99,
		},
	}

	fmt.Println(products)
}
