package main

import (
	"AlgoTrading-Backend/Server"
	"fmt"
	"google.golang.org/appengine"
	"os"
	"log"
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
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	fmt.Printf("Spinning up server on port %s", port)
	fmt.Println("Welcome to AlgoTrading.")
	Server.Routes(port)
	appengine.Main() // Start the server
}
