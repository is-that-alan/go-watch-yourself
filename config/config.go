package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"go-db-project/price"
)

// LoadConfig loads Alpaca API credentials from environment or config file
func LoadConfig() (price.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Path to config subdir
	viper.AutomaticEnv()            // Override with environment variables

	// Default values
	viper.SetDefault("ALPACA_BASE_URL", "https://data.alpaca.markets")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return price.Config{}, fmt.Errorf("failed to read config: %w", err)
		}
		// Config file not found; rely on environment variables
	}

	cfg := price.Config{
		APIKey:    viper.GetString("ALPACA_API_KEY"),
		APISecret: viper.GetString("ALPACA_API_SECRET"),
		BaseURL:   viper.GetString("ALPACA_BASE_URL"),
	}

	if cfg.APIKey == "" || cfg.APISecret == "" {
		return price.Config{}, errors.New("ALPACA_API_KEY and ALPACA_API_SECRET must be set")
	}

	return cfg, nil
}
