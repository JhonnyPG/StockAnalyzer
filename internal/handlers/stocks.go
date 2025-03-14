package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Stock struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	Brokerage  string `json:"brokerage"`
	TargetFrom string `json:"target_from"`
}

func GetStocks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
            SELECT ticker, company, brokerage, target_from 
            FROM stocks
        `)
		if err != nil {
			http.Error(w, "Error en la consulta", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var stocks []Stock
		for rows.Next() {
			var s Stock
			err := rows.Scan(&s.Ticker, &s.Company, &s.Brokerage, &s.TargetFrom)
			if err != nil {
				http.Error(w, "Error al leer datos", http.StatusInternalServerError)
				return
			}
			stocks = append(stocks, s)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stocks)
	}
}
