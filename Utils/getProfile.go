package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type WholeProfile struct {
	Symbol            string	`json:"symbol"`
	Profile struct {
		Price             float64	`json:"price"`
		Beta              string	`json:"beta"`
		VolAvg            string	`json:"volAvg"`
		MktCap            string	`json:"mktCap"`
		LastDiv           string	`json:"lastDiv"`
		Range             string	`json:"range"`
		Changes           float64	`json:"Changes"`
		ChangesPercentage string	`json:"changesPercentage"`
		CompanyName       string	`json:"companyName"`
		Exchange          string	`json:"exchange"`
		Industry          string	`json:"industry"`
		Website           string	`json:"website"`
		Description       string	`json:"description"`
		Ceo               string	`json:"ceo"`
		Sector            string	`json:"sector"`
		Image             string	`json:"image"`
	}	`json:"profile"`
}

var profileUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company/profile/",
}

// Gets profile from financialmodelingprep.com
func GetProfile(symbol string) WholeProfile {
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

	return data

	// return data
}