package Seed

import (
	"github.com/jinzhu/gorm"
	"github.com/m0thm4n/AlgoTrading-Backend/models"
	"log"
)

var users = []models.User {
	models.User {
		Name: 		"Nathan Moritz",
		Email: 		"nathan.moritz@protonmail.com",
		Password:	"P@$$W0rd!",
	},
}