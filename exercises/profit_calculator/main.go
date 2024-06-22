package main

import "fmt"


func main() {
	currentRevenue := inputRequest("What is your current revenue?: ")
	expenses := inputRequest("How much expenses you have?: ")
	taxRate := inputRequest("What is your tax rate?: ")

	earningsBeforeTax, earingsAfterTax, ratio := calculateValues(currentRevenue, expenses, taxRate)

	outputText("EBT: ", earningsBeforeTax)
	outputText("Profit: ", earingsAfterTax)
	outputText("Ratio: ", ratio)
}

func inputRequest(question string) float64 {
	var value float64
	fmt.Print(question)
	fmt.Scan(&value)
	return value
}

func calculateValues(currentRevenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := currentRevenue - expenses
	earingsAfterTax := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / earingsAfterTax

	return earningsBeforeTax, earingsAfterTax, ratio
}

func outputText(text string, value float64) {
	fmt.Print(text)
	fmt.Println(value)
}
