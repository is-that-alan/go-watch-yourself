package cmd

import (
	"encoding/json"
	"fmt"
	"go-watch-yourself/watchlist"
	"os"

	"github.com/spf13/cobra"
	"slices"
)

var removeCmd = &cobra.Command{
	Use:   "remove [symbol]",
	Short: "Remove a stock from your watchlist",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		json_text, err := os.ReadFile("test_cmd.json")
		if err != nil {
			fmt.Errorf("there was an error: %v", err)
		}
		var wl watchlist.WatchList
		err = json.Unmarshal(json_text, &wl)
		if err != nil {
			fmt.Errorf("error while unmarshaling json: %v", err)
		}
		initialCount := len(wl.Items)
		// Remove the item by symbol
		wl.Items = slices.DeleteFunc(wl.Items, func(item watchlist.WatchItem) bool {
			return item.Symbol == symbol
		})

		// Check if anything was actually removed
		if len(wl.Items) == initialCount {
			fmt.Printf("Symbol '%s' not found in watchlist.\n", symbol)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
