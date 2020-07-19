package Utils

import (
	"encoding/json"
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"log"
	"net/http"
	"net/url"
)

type CompanyEnterpriseValue struct {
	Symbol           string `json:"symbol"`
	EnterpriseValues []struct {
		Date                   string  `json:"date"`
		StockPrice             string  `json:"Stock Price"`
		NumberOfShares         string  `json:"Number of Shares"`
		MarketCapitalization   float64 `json:"Market Capitalization"`
		CashAndCashEquivalents float64 `json:"- Cash & Cash Equivalents"`
		TotalDebt              float64 `json:"+ Total Debt"`
		EnterpriseValue        float64 `json:"Enterprise Value"`
	} `json:"enterpriseValues"`
}

var companyEnterpriseValue = url.URL{
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/enterprise-value/",
}

// Series of functions for getting Company Enterprise Value from financialmodelingprep.com
func GetCompanyEnterpriseValue(symbol string) CompanyEnterpriseValue {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(companyEnterpriseValue.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyEnterpriseValue

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCompanyEnterpriseValue(symbol string) CompanyEnterpriseValue {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(companyEnterpriseValue.String() + symbol + "?period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyEnterpriseValue

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
