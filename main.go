package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const API_KEY = "7718D45E-FA3C-4EFD-BDC3-2167FC97B707"

func main() {
    fmt.Println("Enter n to get fib(n):")
    var n int
    fmt.Scanln(&n)
	fmt.Println("fib(n) =", fibonacciRecursion(n))
	fmt.Println("Current BTC price in USD:", getBitcoinPrice())
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

type GetPriceResult struct {
	Time         string  `json:"time"`
	AssetIdBase  string  `json:"asset_id_base"`
	AssetIdQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

func getBitcoinPrice() float64 {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
    if err != nil {
        fmt.Print(err.Error())
		os.Exit(1)
    }
    req.Header.Add("X-CoinAPI-Key", API_KEY)
    response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result GetPriceResult
	json.Unmarshal(responseData, &result)

	return result.Rate
}
