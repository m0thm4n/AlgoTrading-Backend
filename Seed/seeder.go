package Seed

import (
	"github.com/jinzhu/gorm"
	"github.com/m0thm4n/AlgoTrading-Backend/Models"
	"log"
)

var users = []Models.User{
	Models.User{
		Name:     "Nathan Moritz",
		Email:    "nathan.moritz@protonmail.com",
		Password: "P@$$W0rd!",
	},
	Models.User{
		Name:		"Mothman",
		Email:		"nateim3@gmail.com",
		Password: 	"Pa$$w0rD!",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&Models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&Models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&Models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
	}
}
