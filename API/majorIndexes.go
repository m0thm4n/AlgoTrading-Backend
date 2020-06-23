package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMajorIndexes(w http.ResponseWriter, r *http.Request) {
	majorIndexes := Utils.GetMajorIndexes()
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &majorIndexes)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(majorIndexes)
}
