package Server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/m0thm4n/AlgoTrading-Backend/Controllers"
	"github.com/m0thm4n/AlgoTrading-Backend/Seed"
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

	server.Run(":8080")
}
