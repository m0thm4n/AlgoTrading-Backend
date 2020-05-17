package Utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Quote struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Change            float64 `json:"change"`
	DayLow            float64 `json:"dayLow"`
	DayHigh           float64 `json:"dayHigh"`
	YearHigh          float64 `json:"yearHigh"`
	YearLow           float64 `json:"yearLow"`
	MarketCap         float64 `json:"marketCap"`
	PriceAvg50        float64 `json:"priceAvg50"`
	PriceAvg200       float64 `json:"priceAvg200"`
	Volume            int64   `json:"volume"`
	AvgVolume         int64   `json:"avgVolume"`
	Exhange           string  `json:"exhange"`
	Open              float64 `json:"open"`
	PreviousClose     float64 `json:"previousClose"`
	Eps               float64 `json:"eps"`
	Pe                float64 `json:"pe"`
	SharesOutstanding int64   `json:"sharesOutstanding"`
}

type WholeProfile struct {
	Symbol            string  `json:"symbol"`
	Profile struct {
		Price             float64 `json:"price"`
		Beta              string `json:"beta"`
		VolAvg            string `json:"volAvg"`
		MktCap            string `json:"mktCap"`
		LastDiv           string `json:"lastDiv"`
		Range             string `json:"range"`
		Changes           float64 `json:"Changes"`
		ChangesPercentage string `json:"changesPercentage"`
		CompanyName       string  `json:"companyName"`
		Exchange          string  `json:"exchange"`
		Industry          string  `json:"industry"`
		Website           string  `json:"website"`
		Description       string  `json:"description"`
		Ceo               string  `json:"ceo"`
		Sector            string  `json:"sector"`
		Image             string  `json:"image"`
	} `json:"profile"`
}

var quoteUrl = url.URL{
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/quote/",
}

var profileUrl = url.URL{
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company/profile/",
}

func GetQuote(symbol string) []Quote {
	response, err := http.Get(quoteUrl.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []Quote

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetProfile(symbol string) {
	response, err := http.Get(profileUrl.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data WholeProfile

	// err = json.NewDecoder(response.Body).Decode(&data)
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	// return data
}