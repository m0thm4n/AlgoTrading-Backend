package API

import (
	"AlgoTrading-Backend/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHistoricalStockPriceByMinute(w http.ResponseWriter, r *http.Request) {
	time := mux.Vars(r)["time"]
	symbol := mux.Vars(r)["symbol"]
	historicalStockPrice := Utils.GetHistoricalStockPriceByMinute(time, symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &historicalStockPrice)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(historicalStockPrice)
}


func GetHistoricalStockPriceByHour(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	historicalStockPrice := Utils.GetHistoricalStockPriceByHour(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &historicalStockPrice)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(historicalStockPrice)
}

func GetDailyHistoricalStockPrice(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	dailyHistoricalStockPrice := Utils.GetDailyHistoricalStockPrice(symbol)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &dailyHistoricalStockPrice)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dailyHistoricalStockPrice)
}