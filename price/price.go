package price

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"encoding/json"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

// PriceService manages stock price fetching with caching
type PriceService struct {
	client marketdata.Client // Use interface directly, not pointer
	cache  map[string]cachedPrice
	mu     sync.RWMutex
}

// cachedPrice stores a price with its timestamp
type cachedPrice struct {
	price     float64
	timestamp time.Time
}

// Config holds Alpaca API credentials
type Config struct {
	APIKey    string
	APISecret string
	BaseURL   string // Renamed to match SDK's ClientOpts (was Endpoint)
}

// NewPriceService initializes the price service with Alpaca client and cache
func NewPriceService(cfg Config) (*PriceService, error) {
	if cfg.APIKey == "" || cfg.APISecret == "" {
		return nil, errors.New("API key and secret are required")
	}

	client := marketdata.NewClient(marketdata.ClientOpts{
		ApiKey:    cfg.APIKey,
		ApiSecret: cfg.APISecret,
		BaseURL:   cfg.BaseURL,
		Feed:      "iex", // Use free IEX feed to avoid subscription issues
	})

	return &PriceService{
		client: client,
		cache:  make(map[string]cachedPrice),
		mu:     sync.RWMutex{},
	}, nil
}

// GetStockPrice fetches the latest bar's close price for a given ticker
func (s *PriceService) GetStockPrice(ticker string) (float64, error) {
	if ticker == "" {
		return 0, errors.New("ticker cannot be empty")
	}

	// Check cache first
	s.mu.RLock()
	if cached, exists := s.cache[ticker]; exists && time.Since(cached.timestamp) < 1*time.Minute {
		s.mu.RUnlock()
		return cached.price, nil
	}
	s.mu.RUnlock()

	// Fetch from Alpaca API using GetLatestBar for latest data
	bar, err := s.client.GetLatestBar(ticker)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch latest bar for %s: %w", ticker, err)
	}

	if bar == nil || bar.Close == 0 {
		return 0, fmt.Errorf("no bar available for %s", ticker)
	}

	// Update cache with close price
	s.mu.Lock()
	s.cache[ticker] = cachedPrice{
		price:     bar.Close,
		timestamp: time.Now(),
	}
	s.mu.Unlock()

	return bar.Close, nil
}

// SaveMockData captures a bar response as JSON for testing
func (s *PriceService) SaveMockData(ticker, filename string) error {
	bar, err := s.client.GetLatestBar(ticker)
	if err != nil {
		return fmt.Errorf("failed to fetch mock data for %s: %w", ticker, err)
	}

	data, err := json.MarshalIndent(bar, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal mock data: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save mock data to %s: %w", filename, err)
	}

	log.Printf("Mock data saved to %s", filename)
	return nil
}
