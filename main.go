package main

import (
	"AlgoTrading/Utils"
	"AlgoTrading/Server"
	// "fmt"
)

func main() {
	// quote := Utils.GetQuote()
	// fmt.Println(quote)

	// session := Database.DbConnect("localhost:8080/")
	// Database.CreateTable(session)

	Utils.GetTicker("PRAA", "10", "NASDAQ")

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", API.HomeLink)
	// router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8080", router))

	Server.Routes()
}
