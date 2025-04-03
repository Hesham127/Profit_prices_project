package main

import (
	"fmt"
	"price_calculator/filemanger"
	"price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChanels := make([]chan bool, 4)
	errorChanels := make([]chan error, 4)

	for index, taxRate := range taxRates {
		doneChanels[index] = make(chan bool)
		errorChanels[index] = make(chan error)

		fm := filemanger.New("prices.txt", fmt.Sprintf("result_ %0.f", taxRate))
		// cm := cmdmanger.New() if u wanna try care from go rotuines
		priceJob := prices.NewTaxIncludedPriceJop(fm, taxRate)
		go priceJob.Process(doneChanels[index], errorChanels[index])

	}
	for index := range taxRates {

		select {
		case err := <-errorChanels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChanels[index]:
		}
	}

}
