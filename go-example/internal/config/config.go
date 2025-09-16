package config

import (
	"os"
	"fmt"
)

func LoadConfig() {
    fmt.Println("Loading configuration...")
    // Load environment variables or config files here
}


func GetEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}