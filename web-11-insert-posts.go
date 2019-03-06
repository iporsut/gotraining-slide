package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:@/blogdb")
	check(err)

	stmt, err := db.PrepareContext(ctx, "INSERT INTO posts(title, body) values (?,?)")
	check(err)

	result, err := stmt.ExecContext(ctx, "Learn Go", "Today we will learn how to write Go.")
	check(err)
	lastID, _ := result.LastInsertId()
	fmt.Println("New Record ID:", lastID)
}
