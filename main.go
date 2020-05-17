package main

import (
	"AlgoTrading/Utils"
)

func main() {
	// quote := Utils.GetQuote()
	// fmt.Println(quote)

	Utils.GetProfile("AAPL")

	// session := Database.DbConnect("localhost:8080/")
	// Database.CreateTable(session)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", API.HomeLink)
	// router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8080", router))

	// Server.Routes()
}
