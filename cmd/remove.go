package cmd

import (
	"encoding/json"
	"fmt"
	"go-watch-yourself/watchlist"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [symbol]",
	Short: "Remove a stock from your watchlist",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_ = args[0]
		json_text, err := os.ReadFile("test_cmd.json")
		if err != nil {
			fmt.Errorf("there was an error: %v", err)
		}
		var wl watchlist.WatchList
		err = json.Unmarshal(json_text, &wl)
		if err != nil {
			fmt.Errorf("error while unmarshaling json: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
