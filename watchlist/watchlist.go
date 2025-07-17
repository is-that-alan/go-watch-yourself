package watchlist

import (
	"encoding/json"
	"fmt"
	// "log"
	"os"
	"time"
)

const filePath string = "test_json_data.json"

type WatchList struct {
	Items []WatchItem
}

type WatchItem struct {
	Symbol    string    `json:"symbol"`
	Expiry    time.Time `json:"time"`
	Threshold float32   `json:"threshold"`
	Above     bool      `json:"above"`
	IsActive  bool      `json:"isActive"`
}

func AddToWatchList(
	symbol string,
	expiry time.Time,
	threshold float32,
	above bool,
) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("error while reading json file: %v", err)
	}
	var wl WatchList
	err = json.Unmarshal(data, &wl.Items)
	if err != nil {
		fmt.Errorf("error while unmarshaling json file: %v", err)
		return err
	}

	w := WatchItem{
		Symbol:    symbol,
		Expiry:    expiry,
		Threshold: threshold,
		Above:     above,
	}
	wl.Items = append(wl.Items, w)
	jsonData, err := json.Marshal(wl.Items)
	if err != nil {
		fmt.Errorf("error while marshaling data: %v", err)
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Errorf("error while marshaling data: %v", err)
		return err
	}
	return nil

}

// func SaveWatchList() error {
// 	watchList, err := MakeWatchList(
// 		"AAPL",
// 		time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC),
// 		100.0,
// 		true,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	jsonData, err := json.Marshal(watchList)
// 	if err != nil {
// 		log.Fatalf("Error marshaling to JSON: %v", err)
// 	}
// 	err = os.WriteFile("test_json_data.json", jsonData, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
