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
	router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	router.HandleFunc("/api/company/profile/{symbol}", API.CreateProfile).Methods("GET")
	router.HandleFunc("/api/ticker/{query}/{limit}/{exchange}", API.CreateTicker).Methods("GET")
	router.HandleFunc("/api/income/{symbol}", API.CreateIncomeStatment).Methods("GET")

	fmt.Printf("Spinning up server on  %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host + ":" + port, router))
}
