package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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
	db, err := sql.Open("mysql", "root:@/blogdb?parseTime=true")
	check(err)

	// BEGIN QUERY OMIT
	qry := "SELECT id, title, body, created_at, updated_at FROM posts WHERE id = ?"
	stmt, err := db.PrepareContext(ctx, qry)
	check(err)

	rows, err := stmt.QueryContext(ctx, 1)
	check(err)

	for rows.Next() {
		var (
			id        int
			title     string
			body      string
			createdAt time.Time
			updatedAt time.Time
		)
		err := rows.Scan(&id, &title, &body, &createdAt, &updatedAt)
		check(err)
		fmt.Println(id, title, body, createdAt.Format(time.RFC3339), updatedAt.Format(time.RFC3339))
	}
	// END QUERY OMIT

	// BEGIN QUERY ROW OMIT
	qry := "SELECT id, title, body, created_at, updated_at FROM posts WHERE id = ?"
	stmt, err := db.PrepareContext(ctx, qry)
	check(err)

	row := stmt.QueryRowContext(ctx, 1)
	var (
		id        int
		title     string
		body      string
		createdAt time.Time
		updatedAt time.Time
	)
	err = row.Scan(&id, &title, &body, &createdAt, &updatedAt)
	check(err)
	fmt.Println(id, title, body, createdAt.Format(time.RFC3339), updatedAt.Format(time.RFC3339))
	// END QUERY ROW OMIT
}
