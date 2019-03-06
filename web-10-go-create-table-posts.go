package main

import (
	"context"
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

	// BEGIN CREATE TABLE OMIT
	// Execute Create Table
	postsTableDDL := `
CREATE TABLE posts (
       id INT NOT NULL AUTO_INCREMENT,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       title TEXT NOT NULL,
       body TEXT NOT NULL,
       PRIMARY KEY (id)
);
`
	_, err = db.ExecContext(context.Background(), postsTableDDL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// END CREATE TABLE OMIT
}
