package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type CompanyRating struct {
	Symbol	string	`json:"symbol"`
	Rating	[]struct {
		Score				int		`json:"score"`
		Rating				string	`json:"rating"`
		Recommendation		string	`json:"recommendation"`
	} `json:"rating"`
	RatingDetails	[]struct {
		PB []struct {
			Score          int    `json:"score"`
			Recommendation string `json:"recommendation"`
		} `json:"P/B"`
		ROA []struct {
			Score			int    `json:"score"`
			Recommendation	string `json:"recommendation"`
		} `json:"ROA"`
		DCF []struct {
			Score			int		`json:"score"`
			Recommendation	string	`json:"recommendation"`
		} `json:"DCF"`
		PE	[]struct {
			Score			int		`json:"score"`
			Recommendation	string	`json:"recommendation"`
		} `json:"P/E"`
		ROE	[]struct {
			Score			int		`json:"score"`
			Recommendation	string	`json:"recommendation"`
		} `json:"ROE"`
		DE	[]struct {
			Score			int		`json:"score"`
			Recommendation	string	`json:"recommendation"`
		} `json:"D/E"`
	} `json:"ratingDetails"`
}

var companyRatingUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company/rating/",
}

// Gets Company Rating from financialmodelingprep.com
func GetCompanyRating(symbol string) CompanyRating {
	response, err := http.Get(companyRatingUrl.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyRating

	// err = json.NewDecoder(response.Body).Decode(&data)
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data

	// return data
}
