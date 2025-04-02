package prices

import (
	"fmt"
	"price_calculator/conversions"
	"price_calculator/filemanger"
)

type TaxIncludedPriceJop struct {
	IoManger          filemanger.Filemanger `json:"-"`
	TaxRate           float64               `json:"tax_rate"`
	InputPrices       []float64             `json:"input_prices"`
	TaxIncludedPrices map[string]string     `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJop) loadData() {

	lines, err := job.IoManger.ReadFile()
	if err != nil {
		fmt.Println("There is an Error!")
		fmt.Println(err)
		return
	}

	prices, err := conversions.StringsToFloats(lines)

	if err != nil {
		fmt.Println("There is an Error!")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
	// fmt.Printf("loaded Prices %f \n", job.InputPrices)
}
func (job *TaxIncludedPriceJop) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)
	}
	job.TaxIncludedPrices = result
	job.IoManger.WriteResult(job)
}

func NewTaxIncludedPriceJop(fm filemanger.Filemanger, taxRate float64) *TaxIncludedPriceJop {

	return &TaxIncludedPriceJop{
		IoManger:          fm,
		InputPrices:       []float64{10., 20., 30.},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string),
	}
}
