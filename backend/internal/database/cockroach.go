package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/JhonnyPG/StockAnalyzer/internal/models"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func Connect() *sql.DB {
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

	fmt.Println("✅ Conected to CockroachDB " + now.String())
	return db
}

type DBExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func InsertStock(db DBExecutor, stock models.Stock) error {
	query := `
        INSERT INTO stocks 
            (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
        VALUES 
            ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (ticker) DO NOTHING`

	_, err := db.Exec(
		query,
		stock.Ticker,
		stock.Company,
		stock.Brokerage,
		stock.Action,
		stock.RatingFrom,
		stock.RatingTo,
		stock.TargetFrom,
		stock.TargetTo,
		stock.Time,
	)

	if err != nil {
		log.Printf("Error inserting stock %s: %v", stock.Ticker, err)
		return fmt.Errorf("failed to insert stock: %w", err)
	}

	fmt.Printf("✅ Successfully inserted stock: %s\n", stock.Ticker)
	return nil
}
