package main

import (
	"fmt"
	"github.com/m0thm4n/AlgoTrading-Backend/Server"
	// "google.golang.org/appengine"
	// "os"
	// "log"

)


func main() {
	fmt.Println("Welcome to AlgoTrading")
	// quote := Utils.GetQuote()
	// fmt.Println(quote)

	// session := Database.DbConnect("localhost:8080/")
	// Database.CreateTable(session)

	// ticker := Utils.GetTicker("PRAA", "10", "NASDAQ")
	// fmt.Println(ticker)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", API.HomeLink)
	// router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8080", router))

	// fmt.Println("Welcome to AlgoTrading.")
	Server.Run()
}
