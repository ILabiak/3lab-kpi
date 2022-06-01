package main

import (
	"database/sql"
	//forums2 "github.com/ILabiak/3lab-kpi/pkg/forums"
	"github.com/ILabiak/3lab-kpi/pkg/db"
)

func main () {
	
}

func NewDbConnection() (*sql.DB, error) {
	conn := db.Connection{
		DbName:     "forumsdb",
		User:       "postgres",
		Password:   "qwerty334455",
		Host:       "127.0.0.1",
		Port:       "5050",
		DisableSSL: true,
	}
	return conn.Open()
}


