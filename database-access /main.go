package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// to connect to a MySQL DB

var db *sql.DB

func main() {
	// capture connection properties
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "alihussein10X",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "TRIAL",
		AllowNativePasswords: true,
	}

	// get DB handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("DBUSER:", os.Getenv("DBUSER"))
	fmt.Println("DBPASS:", os.Getenv("DBPASS"))
	fmt.Println("Connected")

}
