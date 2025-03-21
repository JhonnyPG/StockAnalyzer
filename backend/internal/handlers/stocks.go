package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JhonnyPG/StockAnalyzer/internal/models"
)

func GetStocks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
            SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time FROM stocks`)

		if err != nil {
			http.Error(w, "Error en la consulta", http.StatusInternalServerError)
			return
		}
		if !rows.Next() {
			http.Error(w, "No hay datos disponibles", http.StatusNotFound)
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
				log.Printf("Error al leer fila: %v", err)
				http.Error(w, "Error al leer datos de la base", http.StatusInternalServerError)
				return
			}
			stocks = append(stocks, s)
			fmt.Printf("stock append: %v\n", s)
		}

		if len(stocks) == 0 {
			http.Error(w, "No hay datos disponibles", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(stocks); err != nil {
			log.Println("Error al codificar la respuesta:", err)
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

func calculateBestStock(db *sql.DB) (models.Stock, error) {
	var bestStock models.Stock
	err := db.QueryRow(`
        SELECT 
            ticker,
            company,
            brokerage,
            action,
            rating_from,
            rating_to,
            target_from,
            target_to,
            time
        FROM stocks 
        ORDER BY (
        REPLACE(target_to, '$', '')::NUMERIC - 
        REPLACE(target_from, '$', '')::NUMERIC
		) DESC 
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
