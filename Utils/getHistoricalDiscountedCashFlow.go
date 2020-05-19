package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type HistoricalDiscountedCashFlow struct {
	Symbol			string	`json:"symbol"`
	HistoricalDCF []struct {
		Date			string	`json:"date"`
		HistoricalDCF	float64	`json:"dcf"`
	}	`json:"historicalDCF"`
}

var historicalDiscountedCashFlow = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company/historical-discounted-cash-flow/",
}

func GetHistoricalDiscountedCashFlow(symbol string) HistoricalDiscountedCashFlow {
	response, err := http.Get(historicalDiscountedCashFlow.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data HistoricalDiscountedCashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyHistoricalDiscountedCashFlow(symbol string) HistoricalDiscountedCashFlow {
	response, err := http.Get(historicalDiscountedCashFlow.String() + symbol + "?period=quarter")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data HistoricalDiscountedCashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}