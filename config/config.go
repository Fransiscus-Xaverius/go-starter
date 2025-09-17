package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
