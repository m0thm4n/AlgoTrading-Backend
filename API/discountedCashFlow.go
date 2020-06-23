package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDiscountedCashFlow(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	discountedCashFlow := Utils.GetDiscountedCashFlow(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &discountedCashFlow)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(discountedCashFlow)
}