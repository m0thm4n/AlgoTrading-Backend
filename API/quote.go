package API

import (
	"encoding/json"
	"fmt"
	"github.com/m0thm4n/AlgoTrading-Backend/Utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	quote := Utils.GetQuote(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &quote)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(quote)
}
