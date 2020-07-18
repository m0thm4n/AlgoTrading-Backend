package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFinancialRatios(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	financialRatios := Utils.GetFinancialRatios(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &financialRatios)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(financialRatios)
}