package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCompanyFinancialGrowth(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyFinancialGrowth := Utils.GetCompanyFinancialGrowth(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyFinancialGrowth)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyFinancialGrowth)
}

func GetQuarterlyCompanyFinancialGrowth(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyFinancialGrowth := Utils.GetQuarterlyCompanyFinancialGrowth(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyFinancialGrowth)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyFinancialGrowth)
}
