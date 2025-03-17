package handlers

import (
	"database/sql"
	"encoding/json"
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

func GetRecommendationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lógica de recomendación (ej: stock con mayor diferencia en target_to - target_from)
		bestStock := calculateBestStock(db)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bestStock)
	}
}
