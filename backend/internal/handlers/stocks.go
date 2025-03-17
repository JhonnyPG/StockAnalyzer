package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JhonnyPG/StockAnalyzer/internal/models"
)

func GetStocks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
            SELECT * FROM stocks
        `)
		if err != nil {
			http.Error(w, "Error en la consulta", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var stocks []models.Stock
		for rows.Next() {
			var s models.Stock
			err := rows.Scan(
				&s.Ticker,
				&s.Company,
				&s.Brokerage,
				&s.Action,
				&s.RatingFrom,
				&s.RatingTo,
				&s.TargetFrom,
				&s.TargetTo,
				&s.Time,
			)
			if err != nil {
				http.Error(w, "Error al leer datos", http.StatusInternalServerError)
				return
			}
			stocks = append(stocks, s)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, "Error post-escaneo: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stocks)
	}
}

func calculateBestStock(db *sql.DB) (models.Stock, error) {
	var bestStock models.Stock
	err := db.QueryRow(`
        SELECT * FROM stocks 
        ORDER BY (target_to::numeric - target_from::numeric) DESC 
        LIMIT 1
    `).Scan(
		&bestStock.Ticker,
		&bestStock.Company,
		&bestStock.Brokerage,
		&bestStock.Action,
		&bestStock.RatingFrom,
		&bestStock.RatingTo,
		&bestStock.TargetFrom,
		&bestStock.TargetTo,
		&bestStock.Time,
	)
	if err != nil {
		return models.Stock{}, fmt.Errorf("error obteniendo recomendación: %w", err)
	}
	return bestStock, nil
}

func GetRecommendationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lógica de recomendación
		bestStock, err := calculateBestStock(db) // Captura el error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bestStock)
	}
}
