package Database

import (
	"encoding/json"
	"fmt"
	"log"

	db "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func DbConnect(url string) *db.Session {
	var err error

	session, err := db.Connect(db.ConnectOpts{
		Address:  url,
		Database: "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	return session
}

func CreateTable(session *db.Session) {
	result, err := db.DB("test").TableCreate("profile").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printStr("*** Create table result: ***")
	printObj(result)
	printStr("\n")
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
