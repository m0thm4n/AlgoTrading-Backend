package Utils

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Ticker struct {
	Symbol				string `json:"symbol"`
	Name				string `json:"name"`
	Currency			string `json:"currency"`
	StockExchange		string `json:"stockExchange"`
	ExchangeShortName	string `json:"exchangeShortName"`
}

var tickerUrl = url.URL {
	Scheme: "https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/",
}

// Get ticker from financialmodelingprep.com
func GetTicker(query, limit, exchange string) []Ticker {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(tickerUrl.String() + "search?query=" + query + "&" + "limit=" + limit + "&" + "exchange=" + exchange + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []Ticker

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}