package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func connect() {
	dsn := "postgresql://trusteduser:gLqBgcUnAVDj2DPsLvoHag@stock-analyzer-9114.j77.aws-us-west-2.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
