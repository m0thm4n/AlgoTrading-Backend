package Database

import (
	db "github.com/rethinkdb/rethinkdb-go"
	"log"
)

func DbConnect(url string) *db.Session {
	var err error

	session, err := db.Connect(db.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return session
}
