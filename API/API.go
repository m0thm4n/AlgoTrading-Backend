package API

import (
	"AlgoTrading/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	quote := Utils.GetQuote(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &quote)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(quote)
}
