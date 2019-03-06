package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:@/blogdb")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Ping to verify a connection
	err = db.Ping()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
