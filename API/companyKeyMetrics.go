package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCompanyKeyMetrics(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyKeyMetrics := Utils.GetCompanyKeyMetrics(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyKeyMetrics)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyKeyMetrics)
}

func GetQuarterlyCompanyKeyMetrics(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyKeyMetrics := Utils.GetQuarterlyCompanyKeyMetrics(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyKeyMetrics)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyKeyMetrics)
}
