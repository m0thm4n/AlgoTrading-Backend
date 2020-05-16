package API

import (
	"AlgoTrading/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

var quote = Utils.GetQuote()

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &quote)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(quote)
}
