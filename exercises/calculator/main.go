package main

import (
	"fmt"
	"danieljonguitud.com/calculator/filemanager"
	"danieljonguitud.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_%.0f.json", taxRate*100))
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go taxIncludedPriceJob.Process(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}
}
