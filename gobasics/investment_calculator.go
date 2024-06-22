package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64 
	var years float64 
	var expectedReturnRate float64 

	outputText("How much you want to invest?: ")
	fmt.Scan(&investmentAmount)

	outputText("For how many years?: ")
	fmt.Scan(&years)

	outputText("What is your expected return rate?: ")
	fmt.Scan(&expectedReturnRate)

	fv, rfv := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFv := fmt.Sprintf("Future value: %.2f\n", fv)

	fmt.Print(formattedFv)
	fmt.Printf("Future real value: %.2f", rfv)
}


func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflationRate / 100, years)

	return fv, rfv
}
