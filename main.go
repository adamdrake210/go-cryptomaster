package main

import (
	"cryptomasters/api"
	"fmt"
	"time"
)

func main() {
	println("Hello, Cryptomasters!")
	go getCryptoInfo("BTC")
	go getCryptoInfo("ETH")
	go getCryptoInfo("BCH")
	time.Sleep(time.Second * 3)
}

func getCryptoInfo(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("Price of %v is $%.2f \n", rate.Currency, rate.Price)
	}
}
