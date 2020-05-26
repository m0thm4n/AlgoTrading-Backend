package Utils

import (
	"AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type MajorIndexes struct {
	MajorIndexesList	struct {
		Ticker		string	`json:"ticker"`
		Changes		float64	`json:"changes"`
		Price		float64	`json:"price"`
		IndexName	string	`json:"indexName"`
	}	`json:"majorIndexesList"`
}

var majorIndexesUrl = url.URL{
	Scheme: "https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/quotes/",
}

func GetMajorIndexes() []MajorIndexes {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(majorIndexesUrl.String() + "index?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []MajorIndexes

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}