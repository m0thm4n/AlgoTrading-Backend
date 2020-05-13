package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"net/url"
	"time"
)

type Profile struct {
	symbol []string `json:"symbol"`
	name string `json:"name"`
	price float32 `json:"price"`
	changesPercentage float32 `json:"changesPercentage"`
	change float32 `json:"change"`
	dayLow float32 `json:"dayLow"`
	dayHigh float32 `json:"dayHigh"`
	yearHigh float32 `json:"yearHigh"`
	yearLow float32 `json:"yearLow"`
	marketCap float32 `json:"marketCap"`
	priceAvg50 float32 `json:"priceAvg50"`
	priceAvg200 float32 `json:"priceAvg200"`
	volume int32 `json:"volume"`
	avgVolume int32 `json:"avgVolume"`
	exhange string `json:"exhange"`
	open float32 `json:"open"`
	previousClose float32 `json:"previousClose"`
	eps float32 `json:"eps"`
	pe float32 `json:"pe"`
	earningsAnnouncement time.Time `json:"earningsAnnouncement"`
	sharesOutstanding int32 `json:"sharesOutstanding"`
	timestamp time.Time `json:"timestamp"`

}

var baseUrl = url.URL{
	Scheme: "https",
	Host: "financialmodelingprep.com",
	Path: "/api/v3/quote/AAPL,FB",
}

func GetProfile() {
	response, err := http.Get(baseUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var profile []Profile

	jsonErr := json.Unmarshal(responseData, &profile)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("%v\n", profile)
}