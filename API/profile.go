package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	profile := Utils.GetProfile(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &profile)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}