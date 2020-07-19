package Controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/m0thm4n/AlgoTrading-Backend/Models"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
	}

	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Println("We are connected to the %s database", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&Models.User{}) // Database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
