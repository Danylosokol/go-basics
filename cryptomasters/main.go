package main

import (
	"fmt"
	"sync"

	"cryptomasters.com/go/api"
)

func main() {
	var currencies = []string{"BTC", "ETH", "BCH"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
