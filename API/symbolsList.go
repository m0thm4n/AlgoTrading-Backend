package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSymbolsList(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	symbolsList := Utils.GetSymbolsList(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &symbolsList)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(symbolsList)
}