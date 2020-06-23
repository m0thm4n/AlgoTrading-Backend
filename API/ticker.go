package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTicker(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["query"]
	limit := mux.Vars(r)["limit"]
	exchange := mux.Vars(r)["exchange"]

	ticker := Utils.GetTicker(query, limit, exchange)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &ticker)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ticker)
}

