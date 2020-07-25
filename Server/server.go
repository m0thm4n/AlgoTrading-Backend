package Server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/m0thm4n/AlgoTrading-Backend/Controllers"
	"github.com/m0thm4n/AlgoTrading-Backend/Seed"
	"google.golang.org/appengine"
	"log"
	"os"
)

var server = Controllers.Server{}

func Run() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	fmt.Println("Got Values")

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	Seed.Load(server.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	fmt.Printf("Spinning up server on port %s", port)
	fmt.Println("Welcome to AlgoTrading.")
	server.Run(":" + port)
	appengine.Main()
	// server.Run(":8080")
}
