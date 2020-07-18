package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHistoricalDiscountedCashFlow(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	historicalDiscountedCashFlow := Utils.GetHistoricalDiscountedCashFlow(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &historicalDiscountedCashFlow)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(historicalDiscountedCashFlow)
}

func GetQuarterlyHistoricalDiscountedCashFlow(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	historicalDiscountedCashFlow := Utils.GetQuarterlyHistoricalDiscountedCashFlow(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &historicalDiscountedCashFlow)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(historicalDiscountedCashFlow)
}
