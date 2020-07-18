package Utils

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type SymbolsList struct {
	SymbolsList []struct {
		Symbol		string	`json:"symbol"`
		Name		string	`json:"name"`
		Price		float64	`json:"price"`
	} `json:"symbolsList"`
}

var symbolsListUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/company/stock/list",
}

func GetSymbolsList(symbol string) SymbolsList {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(symbolsListUrl.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data SymbolsList

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}