package prices

import (
	"fmt"
	"price_calculator/conversions"
	"price_calculator/iomanger"
	"time"
)

type TaxIncludedPriceJop struct {
	IoManger          iomanger.IoManger `json:"-"`
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJop) loadData() error {

	lines, err := job.IoManger.ReadFile()
	if err != nil {
		return err
	}

	prices, err := conversions.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	// fmt.Printf("loaded Prices %f \n", job.InputPrices)
	return nil
}
func (job *TaxIncludedPriceJop) Process(doneChanel chan bool, errorChanel chan error) error {
	err := job.loadData()
	if err != nil {
		errorChanel <- err
		return err
	}
	time.Sleep(3 * time.Second)
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)
	}
	job.TaxIncludedPrices = result
	job.IoManger.WriteResult(job)
	fmt.Printf("DONE! %f ", job.TaxRate)
	doneChanel <- true
	return nil
}

func NewTaxIncludedPriceJop(ioManger iomanger.IoManger, taxRate float64) *TaxIncludedPriceJop {

	return &TaxIncludedPriceJop{
		IoManger:          ioManger,
		InputPrices:       []float64{10., 20., 30.},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string),
	}
}
