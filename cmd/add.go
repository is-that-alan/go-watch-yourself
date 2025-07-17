package cmd

import (
	// "encoding/json"
	"fmt"
	// "os"
	"time"

	"go-watch-yourself/watchlist"

	"github.com/spf13/cobra"
)

var (
	threshold float32
	expiryStr string
	above     bool
)

var setCmd = &cobra.Command{
	Use:   "set [symbol]",
	Short: "Set a price alert for a stock",
	Args:  cobra.ExactArgs(1), // 1 required positional arg
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		var expiry time.Time
		var err error
		if expiryStr != "" {
			expiry, err = time.Parse("2006-01-02", expiryStr)
			if err != nil {
				return fmt.Errorf("invalid expiry date: %v", err)
			}
		} else {
			expiry = time.Now().Add(30 * 24 * time.Hour)
		}

		err = watchlist.AddToWatchList(symbol, expiry, threshold, above)
		if err != nil {
			return err
		}
		// jsonData, err := json.Marshal(watch)
		// if err != nil {
		// 	return fmt.Errorf("json marshal error: %v", err)
		// }
		// os.WriteFile("test_cmd.json", jsonData, 0644)
		// fmt.Printf("Created watchlist entry: %+v\n", watch)
		return nil
	},
}

func init() {
	setCmd.Flags().Float32VarP(&threshold, "threshold", "t", 0.0, "Price threshold")
	setCmd.Flags().BoolVar(&above, "above", false, "Trigger when price goes above threshold")
	setCmd.Flags().StringVar(&expiryStr, "expiry", "", "Expiry date (YYYY-MM-DD)")

	setCmd.MarkFlagRequired("threshold")

	rootCmd.AddCommand(setCmd)
}
