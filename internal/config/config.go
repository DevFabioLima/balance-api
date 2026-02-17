package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port              string
	InfuraURL         string
	RequestTimeout    time.Duration
	HistoryListLimit  int
	HistoryMaxRecords int
}

func Load() (Config, error) {
	cfg := Config{
		Port:              getEnvOrDefault("PORT", "8080"),
		InfuraURL:         os.Getenv("INFURA_URL"),
		RequestTimeout:    time.Duration(getIntEnvOrDefault("REQUEST_TIMEOUT_MS", 10000)) * time.Millisecond,
		HistoryListLimit:  getIntEnvOrDefault("HISTORY_LIST_LIMIT", 20),
		HistoryMaxRecords: getIntEnvOrDefault("HISTORY_MAX_RECORDS", 1000),
	}

	if cfg.InfuraURL == "" {
		return Config{}, fmt.Errorf("INFURA_URL is required")
	}
	if cfg.HistoryListLimit <= 0 {
		return Config{}, fmt.Errorf("HISTORY_LIST_LIMIT must be > 0")
	}
	if cfg.HistoryMaxRecords <= 0 {
		return Config{}, fmt.Errorf("HISTORY_MAX_RECORDS must be > 0")
	}
	return cfg, nil
}

func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getIntEnvOrDefault(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
