package main

import (
	"AlgoTrading/API"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// quote := Utils.GetQuote()
	// fmt.Println(quote)

	// session := Database.DbConnect("localhost:8080/")
	// Database.CreateTable(session)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", API.HomeLink)
	router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
