package API

import (
	"AlgoTrading/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mux "github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	quote := Utils.GetQuote(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &quote)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(quote)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	profile := Utils.GetProfile(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &profile)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}

func CreateTicker(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["query"]
	limit := mux.Vars(r)["limit"]
	exchange := mux.Vars(r)["exchange"]

	ticker := Utils.GetTicker(query, limit, exchange)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &ticker)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ticker)
}

func CreateIncomeStatment(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	incomeStatment := Utils.GetIncomeStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &incomeStatment)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(incomeStatment)
}