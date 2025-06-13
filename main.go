package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// Query the "users" table
	rows, err := conn.Query(context.Background(),
		`
		SELECT 
			t1.employeeid, t1.firstname, t1.companyid, t2.companyname 
		FROM employee t1
		LEFT JOIN company t2 ON t1.companyid = t2.companyid 
	;`)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	// Iterate and print results
	for rows.Next() {
		var employeeid int
		var firstname string
		var companyid int
		var companyname string

		err := rows.Scan(&employeeid, &firstname, &companyid, &companyname)
		if err != nil {
			log.Printf("Row scan failed: %v\n", err)
			continue
		}

		fmt.Printf("employeeid: %d, firstname: %s, companyid: %d, company name=%s\n", employeeid, firstname, companyid, companyname)
	}

	// Check for errors after loop
	if rows.Err() != nil {
		log.Fatalf("Error during rows iteration: %v\n", rows.Err())
	}
}
