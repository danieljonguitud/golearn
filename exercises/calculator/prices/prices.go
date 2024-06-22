package prices

import (
	"strconv"
	"time"
	"danieljonguitud.com/calculator/conversion"
	"danieljonguitud.com/calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate float64 `json:"taxRate"`
	InputPrices []float64 `json:"inputPrices"`
	TaxIncludedPrices map[string]string `json:"taxIncludedPrices"`
	IOManager iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}	

	prices, err := conversion.StringsToFloats(lines)

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	time.Sleep(3 * time.Second)

	if err != nil {
		errorChan <- err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		value := price * (1 + job.TaxRate)
		result[strconv.FormatFloat(price, 'f', 2, 64)] = strconv.FormatFloat(value, 'f', 2, 64)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteData(job)

	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate: taxRate,
		IOManager: iom,
	}
}
