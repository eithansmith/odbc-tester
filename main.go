package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	dsn, uid, pwd, err := setVars()
	if err != nil {
		log.Fatal(err)
	}

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

	//goland:noinspection SqlNoDataSourceInspection,SqlResolve
	query := `SELECT COUNT(*) FROM PCIDLIB.CIMASTRN`

	var count int64
	err = db2.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Count: %d\n", count)
}

func setVars() (string, string, string, error) {
	var dsn, uid, pwd string
	var err error

	dsn, err = env("DB2_DSN")
	if err != nil {
		return "", "", "", err
	}

	uid, err = env("DB2_UID")
	if err != nil {
		return "", "", "", err
	}

	pwd, err = env("DB2_PWD")
	if err != nil {
		return "", "", "", err
	}

	return dsn, uid, pwd, nil
}

func env(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, fmt.Errorf("%s not set", key)
	}
	return v, nil
}
