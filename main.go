package main

import (
	"AlgoTrading/Database"
	"AlgoTrading/Utils"
)

func main() {
	quote := Utils.GetProfile()

	Database.DbConnect("localhost")
}

