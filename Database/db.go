package Database

import (
	"AlgoTrading/Utils"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Db(quote []Utils.Quote) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&quote)

	// Create
	db.Create(&quote)
}