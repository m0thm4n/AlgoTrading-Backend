package API

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCompanyRating(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	companyRating := Utils.GetCompanyRating(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &companyRating)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyRating)
}