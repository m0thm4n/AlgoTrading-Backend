package Utils

import (
	"AlgoTrading-Backend/Config"
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

type DailyHistoricalStockPrice struct {
	Symbol			string	`json:"symbol"`
	Historical		[]struct {
		Date					string	`json:"date"`
		Open					float64	`json:"open"`
		High					float64	`json:"high"`
		Low						float64	`json:"low"`
		Close					float64	`json:"close"`
		Volume					float64	`json:"volume"`
		UnadjustedVolume		float64	`json:"unadjustedVolume"`
		Change					float64	`json:"change"`
		ChangePercent			float64	`json:"changePercent"`
		VWAP					float64	`json:"vwap"`
		Label					string	`json:"label"`
		ChangeOverTime			float64	`json:"changeOverTime"`
	} `json:"historical"`
}

var historicalStockPriceUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/historical-chart/",
}

var dailyHistoricalStockPriceUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/historical-price-full/",
}

func GetHistoricalStockPrice(time, symbol string) []HistoricalStockPrice {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(historicalStockPriceUrl.String() + time + "min" + "/" + symbol + "?apikey=" + key)
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

func GetDailyHistoricalStockPrice(symbol string) DailyHistoricalStockPrice {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(dailyHistoricalStockPriceUrl.String() + symbol + "?serietype=line&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data DailyHistoricalStockPrice

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}