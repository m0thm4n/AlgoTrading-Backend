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
	router.HandleFunc("/api/quote/{symbol}", API.GetQuote).Methods("GET")
	router.HandleFunc("/api/company/profile/{symbol}", API.GetProfile).Methods("GET")
	router.HandleFunc("/api/ticker/{query}/{limit}/{exchange}", API.GetTicker).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/{symbol}", API.GetIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/quarterly/{symbol}", API.GetQuarterlyIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/csv/{symbol}", API.GetCsvIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet-statement/{symbol}", API.GetBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet/quarterly/{symbol}", API.GetQuarterlyBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet/csv/{symbol}", API.GetCsvBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet/quarterly/csv/{symbol}", API.GetQuarterlyCsvBalanceSheet).Methods("GET")

	fmt.Printf("Spinning up server on  %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host + ":" + port, router))
}
