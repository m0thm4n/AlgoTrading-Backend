package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type CompaniesPriceList struct {
	StockList []struct {
		Symbol		string	`json:"symbol"`
		Price		float64	`json:"price"`
	}	`json:"stockList"`
}

type CompanyPrice struct {
	Symbol		string	`json:"symbol"`
	Price		float64	`json:"price"`
}

var realTimeCompanyStockPrices = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/stock/real-time-price/",
}

var realTimeCompanyStockPriceForCompany = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/stock/real-time-price/AAPL",
}

func GetStockPrices() CompaniesPriceList {
	response, err := http.Get(realTimeCompanyStockPrices.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompaniesPriceList

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetStockPriceForCompany(symbol string) CompanyPrice {
	response, err := http.Get(realTimeCompanyStockPriceForCompany.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyPrice

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}