package Utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Quote struct {
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Change float64 `json:"change"`
	DayLow float64 `json:"dayLow"`
	DayHigh float64 `json:"dayHigh"`
	YearHigh float64 `json:"yearHigh"`
	YearLow float64 `json:"yearLow"`
	MarketCap float64 `json:"marketCap"`
	PriceAvg50 float64 `json:"priceAvg50"`
	PriceAvg200 float64 `json:"priceAvg200"`
	Volume int64 `json:"volume"`
	AvgVolume int64 `json:"avgVolume"`
	Exhange string `json:"exhange"`
	Open float64 `json:"open"`
	PreviousClose float64 `json:"previousClose"`
	Eps float64 `json:"eps"`
	Pe float64 `json:"pe"`
	SharesOutstanding int64 `json:"sharesOutstanding"`
}

type Profile struct {
	Symbol string `json:"symbol"`
	Price float64 `json:"price"`
	Beta float64 `json:"beta"`
	VolAvg float64 `json:"volAvg"`
	MktCap float64 `json:"mktCap"`
	LastDiv float32 `json:"lastDiv"`
	Range float64 `json:"range"`
	Changes float32 `json:"Changes"`
	ChangesPercentage float32 `json:"changesPercentage"`
	CompanyName string `json:"companyName"`
	Exchange string `json:"exchange"`
	Industry string `json:"industry"`
	Website string `json:"website"`
	Description string `json:"description"`
	Ceo string `json:"ceo"`
	Sector string `json:"sector"`
	Image string `json:"image"`
}

var quoteUrl = url.URL{
	Scheme: "https",
	Host: "financialmodelingprep.com",
	Path: "/api/v3/quote/",
}

var profileUrl = url.URL{
	Scheme:     "https",
	Host:       "financialmodelingprep.com",
	Path:       "/api/v3/company/profile",
}

func GetProfile() {
	response, err := http.Get(quoteUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []Quote

	jsonErr := json.NewDecoder(response.Body).Decode(&data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(data)
}