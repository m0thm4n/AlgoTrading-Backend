package Utils

import (
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
	}
}

var symbolsListUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:	"/api/v3/company/stock/list",
}

func GetSymbolsList(symbol string) SymbolsList {
	response, err := http.Get(symbolsListUrl.String() + symbol)
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