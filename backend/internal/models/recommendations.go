package models

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ratingWeights = map[string]int{
	"Strong Sell": 0,
	"Sell":        1,
	"Neutral":     2,
	"Hold":        3,
	"Buy":         4,
	"Strong Buy":  5,
}

type Recommendation struct {
	Ticker      string  `json:"ticker"`
	Company     string  `json:"company"`
	Score       float64 `json:"score"`
	TargetFrom  string  `json:"target_from"`
	TargetTo    string  `json:"target_to"`
	RatingFrom  string  `json:"rating_from"`
	RatingTo    string  `json:"rating_to"`
	LastUpdated string  `json:"last_updated"`
}

func parsePrice(priceStr string) float64 {
	cleanStr := strings.ReplaceAll(priceStr, "$", "")
	cleanStr = strings.ReplaceAll(cleanStr, ",", "")

	price, err := strconv.ParseFloat(cleanStr, 64)
	if err != nil {
		log.Printf("Error parsing price: %s - %v", priceStr, err)
		return 0.0
	}
	return price
}

func calculateScore(stock Stock) float64 {
	// Validar ratings
	fromWeight, ok1 := ratingWeights[stock.RatingFrom]
	toWeight, ok2 := ratingWeights[stock.RatingTo]
	if !ok1 || !ok2 {
		log.Printf("Invalid rating detected: From=%s, To=%s", stock.RatingFrom, stock.RatingTo)
		return 0
	}

	// Calcular componentes
	ratingDiff := toWeight - fromWeight
	fromPrice := parsePrice(stock.TargetFrom)
	toPrice := parsePrice(stock.TargetTo)

	priceChange := 0.0
	if fromPrice > 0 {
		priceChange = (toPrice - fromPrice) / fromPrice * 100
	}

	stockTime := stock.Time

	hoursAgo := time.Since(stockTime).Hours()
	recencyScore := math.Max(0, 100-(hoursAgo/24)*10)

	return (priceChange * 0.5) + (float64(ratingDiff) * 20) + recencyScore
}

func GetRecommendations(stocks []Stock) []Recommendation {
	var recs []Recommendation

	for _, stock := range stocks {
		score := calculateScore(stock)
		if score <= 0 {
			continue
		}

		stockTime := stock.Time

		recs = append(recs, Recommendation{
			Ticker:      stock.Ticker,
			Company:     stock.Company,
			Score:       math.Round(score*100) / 100,
			TargetFrom:  stock.TargetFrom,
			TargetTo:    stock.TargetTo,
			RatingFrom:  stock.RatingFrom,
			RatingTo:    stock.RatingTo,
			LastUpdated: fmt.Sprintf("%.0f hours ago", time.Since(stockTime).Hours()),
		})
	}

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Score > recs[j].Score
	})

	if len(recs) > 5 {
		return recs[:5]
	}
	return recs
}
