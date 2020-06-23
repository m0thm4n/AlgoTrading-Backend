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

type MajoreIndexesQuote struct {
	Symbol					string	`json:"symbol"`
	Name					string	`json:"name"`
	Price					float64	`json:"price"`
	ChangesPercentage		float64	`json:"changesPercentage"`
	Change					float64	`json:"change"`
	DayLow					float64	`json:"dayLow"`
	DayHigh					float64	`json:"dayHigh"`
	YearHigh				float64	`json:"yearHigh"`
	YearLow					float64	`json:"yearLow"`
	PriceAvg50				float64	`json:"priceAvg50"`
	PriceAvg200				float64	`json:"priceAvg200"`
	Volume					int32	`json:"volume"`
	AvgVolume				int64	`json:"avgVolume"`
	Exchange				string	`json:"exchange"`
	Open					float64	`json:"open"`
	PreviousClose			float64	`json:"previousClose"`
	Timestamp				int64	`json:"timestamp"`
}

var majorIndexesUrl = url.URL{
	Scheme: "https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/major-indexes",
}

var majorIndexesQuotesUrl = url.URL {
	Scheme: "https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/quotes/",
}

func GetMajorIndexes() MajorIndexes {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(majorIndexesUrl.String() + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data MajorIndexes

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}