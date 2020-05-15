package main

import (
	"AlgoTrading/Database"
	"AlgoTrading/Utils"
	"fmt"
)

func main() {
	quote := Utils.GetQuote()
	fmt.Println(quote)

	profile := Utils.GetProfile()
	fmt.Println(profile)

	session := Database.DbConnect("localhost:8080/")
	Database.CreateTable(session)
}
