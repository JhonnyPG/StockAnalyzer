package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JhonnyPG/StockAnalyzer/pkg/api"
)

const (
	apiURL    = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MSwiZW1haWwiOiJqaG9ubnkucGVyZWFAb3V0bG9vay5jb20iLCJleHAiOjE3NDIwMDgwNDAsImlkIjoiMCIsInBhc3N3b3JkIjoiJyBPUiAnMSc9JzEifQ.JjFG0jSjU4iN-4YbF4Qvz9R6lbsxG6wpNeM9nRRPUP8"
)

func main() {

	//Create an Api Client
	client := api.NewClient(apiURL, authToken)

	resp, err := client.GetStocks("")
	if err != nil {
		log.Fatalf("Error Getting stocks: %v", err)
	}

	// show obtained data
	fmt.Printf("Retrieved %d stocks\n", len(resp.Items))
	for i, item := range resp.Items {
		fmt.Printf("%d. %s (%s) - %s by %s - Rating: %s -> %s, Target: %s -> %s\n",
			i+1, item.Ticker, item.Company, item.Action, item.Brokerage,
			item.RatingFrom, item.RatingTo, item.TargetFrom, item.TargetTo)
	}

	fmt.Printf("Next page: %s\n", resp.NextPage)

	if len(os.Args) > 1 && os.Args[1] == "all" {
		allStock, err := client.GetAllStock()
		if err != nil {
			log.Fatalf("error getting all stock: %v", err)
		}
		fmt.Printf("\nRetrieved a total of %d stocks\n", len(allStock))
	}

}
