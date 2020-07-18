package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetIncomeStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	incomeStatement := Utils.GetIncomeStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &incomeStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(incomeStatement)
}

func GetQuarterlyIncomeStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	incomeStatement := Utils.GetQuarterlyIncomeStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &incomeStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(incomeStatement)
}

func GetCsvIncomeStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	incomeStatement := Utils.GetQuarterlyIncomeStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &incomeStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(incomeStatement)
}

func GetQuarterlyCsvIncomeStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	incomeStatement := Utils.GetQuarterlyIncomeStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &incomeStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(incomeStatement)
}