package main

import (
	"log"
	"net/http"

	"github.com/JhonnyPG/StockAnalyzer/internal/database"
	"github.com/JhonnyPG/StockAnalyzer/internal/handlers"
)

func main() {
	db := database.Connect()
	defer db.Close()

	http.HandleFunc("/stocks", handlers.GetStocks(db))
	http.HandleFunc("/recommendations", handlers.GetRecommendationHandler(db))

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
