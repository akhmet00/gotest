package db

import (
	"database/sql"
	"fmt"
	"gotest/model"
	"log"

	_ "github.com/lib/pq"
)

func OpenDBConnection() *sql.DB {

	config := model.ReadConfig()

	connString:= fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.Name, config.DbName, config.Pass, config.Address, config.SSLMode)

	db, err := sql.Open("postgres", connString)

	if err!=nil {
		panic(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err!=nil {
		panic(err)
	}
	log.Print("Connected")

	return db
}
