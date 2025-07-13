package main

import (
	"log"

	"go-db-project/config"
	"go-db-project/price"
	"go-db-project/watchlist"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize price service
	svc, err := price.NewPriceService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize price service: %v", err)
	}

	// Fetch real-time price for AAPL
	price, err := svc.GetStockPrice("AAPL")
	if err != nil {
		log.Fatalf("Failed to get price: %v", err)
	}
	log.Printf("AAPL Price: $%.2f", price)

	// Save mock data for testing
	err = svc.SaveMockData("AAPL", "mock_aapl_quote.json")
	if err != nil {
		log.Fatalf("Failed to save mock data: %v", err)
	}

	err = watchlist.SaveWatchList()
	if err != nil {
		log.Fatalf("Failed to save mock watch list data: %v", err)
	}

	log.Println("Save to Json")

}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"sync"
// 	"time"
// )

// type StockPrice struct {
// 	Symbol    string
// 	Price     float64
// 	Timestamp time.Time
// }

// type AlertRule struct {
// 	ID        int
// 	Symbol    string
// 	Threshold float64
// 	Above     bool
// }

// type Alert struct {
// 	RuleID    int
// 	Symbol    string
// 	Threshold float64
// 	Timestamp time.Time
// }

// // TODO add actual implmentation of stock data here
// func generatePriceUpdates(symbol string, priceChan chan<- StockPrice, stopChan <-chan struct{}) {
// 	ticker := time.NewTicker(1 * time.Second) // Simulate price every second
// 	defer ticker.Stop()

// 	basePrice := 100.0 // Starting price
// 	for {
// 		select {
// 		case <-stopChan:
// 			return
// 		case <-ticker.C:
// 			// Simulate price fluctuation
// 			price := basePrice + (rand.Float64()-0.5)*10 // Random walk
// 			priceChan <- StockPrice{
// 				Symbol:    symbol,
// 				Price:     price,
// 				Timestamp: time.Now(),
// 			}
// 			basePrice = price
// 		}
// 	}
// }
