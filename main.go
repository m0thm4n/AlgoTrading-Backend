package main

import (
	"AlgoTrading/API"
	"AlgoTrading/Utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	quote := Utils.GetQuote()
	fmt.Println(quote)

	// Database.DbConnect("localhost")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", API.HomeLink)
	router.HandleFunc("/event", API.CreateQuote)
	log.Fatal(http.ListenAndServe(":8080", router))
}
