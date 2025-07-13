package watchlist

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type WatchList struct {
	Symbol    string    `json:"symbol"`
	Expiry    time.Time `json:"time"`
	Threshold float32   `json:"threshold"`
	Above     bool      `json:"above"`
	IsActive  bool      `json:"isActive"`
}

func MakeWatchList(
	symbol string,
	expiry time.Time,
	threshold float32,
	above bool,
	isActive bool,
) (*WatchList, error) {
	return &WatchList{
		Symbol:    symbol,
		Expiry:    expiry,
		Threshold: threshold,
		Above:     above,
		IsActive:  isActive,
	}, nil
}

func SaveWatchList() error {
	watchList, err := MakeWatchList(
		"AAPL",
		time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC),
		100.0,
		true,
		true,
	)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(watchList)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}
	err = os.WriteFile("test_json_data.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
