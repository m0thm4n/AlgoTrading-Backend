package Server

import (
	"AlgoTrading/API"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	port := "8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", API.HomeLink)
	router.HandleFunc("/api/quote/{symbol}", API.CreateQuote).Methods("GET")
	router.HandleFunc("/api/company/profile/{symbol}", API.CreateProfile).Methods("GET")
	fmt.Printf("Spinning up server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
