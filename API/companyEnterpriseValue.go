package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCompanyEnterpriseValue(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyEnterpriseValue := Utils.GetCompanyEnterpriseValue(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyEnterpriseValue)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyEnterpriseValue)
}

func GetQuarterlyCompanyEnterpriseValue(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyEnterpriseValue := Utils.GetQuarterlyCompanyEnterpriseValue(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyEnterpriseValue)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyEnterpriseValue)
}
