package cmdmanger

import (
	"fmt"
)

type CMDmanger struct {
}

func (CMDm CMDmanger) ReadFile() (prices []string, err error) {
	fmt.Println("PLease enter your prices separeted by ENTER! (if u wanna Exit click 0 : ")
	for {
		var price string
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	// there is no way here to return error

	return prices, nil

}

func (CMDm CMDmanger) WriteResult(data any) error {
	_, err := fmt.Println(data)
	if err != nil {
		return err
	}

	return nil
}
func New() CMDmanger {
	return CMDmanger{}
}
