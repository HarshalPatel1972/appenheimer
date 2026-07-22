package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
	Env  string
}

func Load() (*Config, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT environment variable: %w", err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		Port: port,
		Env:  env,
	}, nil
}
