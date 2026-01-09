package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/alexbrainman/odbc"
)

const (
	dsn = "DEV"
	uid = "user-id-here"
	pwd = "password-here"
)

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	conn := fmt.Sprintf("DSN=%s;UID=%s;PWD=%s", dsn, uid, pwd)
	db2, err := sql.Open("odbc", conn)
	if err != nil {
		log.Fatal(err)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer db2.Close()

	db2.SetMaxOpenConns(1)
	db2.SetMaxIdleConns(1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err := db2.PingContext(ctx); err != nil {
		fmt.Printf("PING ERR: %#v\n", err)
		fmt.Printf("PING ERR (%%v): %v\n", err)
		fmt.Printf("PING ERR (%%+v): %+v\n", err)
		log.Fatalf("DB2 ping failed: %v", err)
	}

	//goland:noinspection SqlNoDataSourceInspection
	query := `SELECT COUNT(*) FROM PCIDLIB.CIMASTRN`

	rows, err := db2.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	var count int64

	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Count: %d\n", count)
}
