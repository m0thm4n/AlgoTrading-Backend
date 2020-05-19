package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type HistoricalStockPrice struct {
	Date		string	`json:"date"`
	Open		float64	`json:"open"`
	Low			float64	`json:"low"`
	High		float64	`json:"high"`
	Close		float64	`json:"close"`
	Volume		float64	`json:"volume"`
}

var stockHistoricalPriceUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/historical-chart/",
}

func GetHistoricalStockPrice(time, symbol string) []HistoricalStockPrice {
	response, err := http.Get(stockHistoricalPriceUrl.String() + time + "min" + "/" + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []HistoricalStockPrice

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
