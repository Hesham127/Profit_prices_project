package main

import (
	"fmt"
	"price_calculator/filemanger"
	"price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanger.New("prices.txt", fmt.Sprintf("result_ %0.f", taxRate))
		priceJob := prices.NewTaxIncludedPriceJop(fm, taxRate)
		priceJob.Process()
	}

}
