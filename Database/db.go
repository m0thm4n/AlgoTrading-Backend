package Database

import (
	"log"

	db "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func DbConnect(url string) {
	var err error
	
	session, err = db.Connect(db.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatalln(err)
	}
}