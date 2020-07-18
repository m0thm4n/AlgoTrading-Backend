package Utils

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
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
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(historicalDiscountedCashFlow.String() + symbol + "?apikey=" + key)
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
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(historicalDiscountedCashFlow.String() + symbol + "?period=quarter&apikey=" + key)
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