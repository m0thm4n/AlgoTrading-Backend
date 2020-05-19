package Utils

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

