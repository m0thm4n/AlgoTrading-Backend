package Server

import (
	"AlgoTrading/API"
	"AlgoTrading/Config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	config := Config.LoadConfiguration("config.json")
	host := config.Host
	port := config.Port

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", API.HomeLink)
	router.HandleFunc("/api/quote/{symbol}", API.GetQuote).Methods("GET")
	router.HandleFunc("/api/company/profile/{symbol}", API.GetProfile).Methods("GET")
	router.HandleFunc("/api/ticker/{query}/{limit}/{exchange}", API.GetTicker).Methods("GET")
	router.HandleFunc("/api/financials/income/{symbol}", API.GetIncomeStatment).Methods("GET")
	router.HandleFunc("/api/financials/income/quarterly/{symbol}", API.GetQuarterlyIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income/csv/{symbol}", API.GetCsvIncomeStatement).Methods("GET")
	router.HandleFunc("/api/income/quarterly/csv/{symbol}", API.GetQuarterlyCsvIncomeStatement).Methods("GET")

	fmt.Printf("Spinning up server on  %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host + ":" + port, router))
}
