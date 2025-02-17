package api

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lucabockmann/go_boilerplate/helper"
)

var db *sqlx.DB

func ConnectDatabase() *sqlx.DB {
	db_user, err := helper.GetEnvDataByName("DBUSER")
	if err != nil {
		log.Fatal(err)
	}

	db_pass, err := helper.GetEnvDataByName("DBPASS")
	if err != nil {
		log.Fatal(err)
	}

	db_url, err := helper.GetEnvDataByName("DBURL")
	if err != nil {
		log.Fatal(err)
	}

	db_port, err := helper.GetEnvDataByName("DBPORT")
	if err != nil {
		log.Fatal(err)
	}

	db_name, err := helper.GetEnvDataByName("DBNAME")
	if err != nil {
		log.Fatal(err)
	}

	db, err = sqlx.Connect("mysql", db_user+":"+db_pass+"@tcp("+db_url+":"+db_port+")"+"/"+db_name)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
