package Utils

import (
	"AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type DiscountedCashFlow struct {
	Symbol		string	`json:"symbol"`
	Date		string	`json:"date"`
	DFC			float64	`json:"dfc"`
	StockPrice	float64	`json:"Stock Price"`
}

var discountedCashFlow = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company/discounted-cash-flow/",
}


// Series of functions for getting Discounted Cash Flow and Historical Discounted Cash Flow from financialmodelingprep.com
func GetDiscountedCashFlow(symbol string) DiscountedCashFlow {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(discountedCashFlow.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data DiscountedCashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}