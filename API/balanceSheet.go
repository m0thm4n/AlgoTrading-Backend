package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBalanceSheet(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	balanceSheet := Utils.GetBalanceSheet(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &balanceSheet)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(balanceSheet)
}

func GetQuarterlyBalanceSheet(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	balanceSheet := Utils.GetQuarterlyBalanceSheet(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &balanceSheet)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(balanceSheet)
}

func GetCsvBalanceSheet(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	balanceSheet := Utils.GetQuarterlyBalanceSheet(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &balanceSheet)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(balanceSheet)
}

func GetQuarterlyCsvBalanceSheet(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	balanceSheet := Utils.GetQuarterlyBalanceSheet(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &balanceSheet)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(balanceSheet)
}