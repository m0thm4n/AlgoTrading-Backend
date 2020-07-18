package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRealTimePrices(w http.ResponseWriter, r *http.Request) {
	// symbol := mux.Vars(r)["symbol"]
	stockPrices := Utils.GetStockPrices()
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &stockPrices)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(stockPrices)
}

func GetRealTimePriceForCompany(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	stockPrice := Utils.GetStockPriceForCompany(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &stockPrice)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(stockPrice)
}

