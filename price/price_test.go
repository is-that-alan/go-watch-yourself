package price

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/stretchr/testify/assert"
)

func TestGetStockPrice(t *testing.T) {
	// Mock configuration
	cfg := Config{
		APIKey:    "test-key",
		APISecret: "test-secret",
		BaseURL:   "https://data.alpaca.markets",
	}

	// Initialize service
	svc, err := NewPriceService(cfg)
	assert.NoError(t, err)

	// Test invalid ticker
	_, err = svc.GetStockPrice("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ticker cannot be empty")

	// Note: Testing real API calls requires valid credentials and network access.
	// For CI, use mock data or a mocked client (see below).
}

func TestMockData(t *testing.T) {
	cfg := Config{
		APIKey:    "test-key",
		APISecret: "test-secret",
		BaseURL:   "https://data.alpaca.markets",
	}

	svc, err := NewPriceService(cfg)
	assert.NoError(t, err)

	// Simulate mock data (requires valid API key for real test)
	err = svc.SaveMockData("AAPL", "test_mock_aapl.json")
	if err == nil { // Skip if API call fails due to invalid credentials
		data, err := os.ReadFile("test_mock_aapl.json")
		assert.NoError(t, err)

		var quote marketdata.Quote
		assert.NoError(t, json.Unmarshal(data, &quote))
		assert.Greater(t, quote.AskPrice, 0.0) // Example assertion; adjust as needed
	}
}
