package API

import (
	"encoding/json"
	"fmt"
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCashFlowStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	cashFlowStatement := Utils.GetCashFlowStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &cashFlowStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cashFlowStatement)
}

func GetQuarterlyCashFlowStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	cashFlowStatement := Utils.GetQuarterlyCashFlowStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &cashFlowStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cashFlowStatement)
}

func GetCsvCashFlowStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	cashFlowStatement := Utils.GetQuarterlyCashFlowStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &cashFlowStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cashFlowStatement)
}

func GetQuarterlyCsvCashFlowStatement(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	cashFlowStatement := Utils.GetQuarterlyCashFlowStatement(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &cashFlowStatement)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cashFlowStatement)
}
