package config

import "os"

type Config struct {
	HTTPAddr string
}

func Load() Config {
	return Config{
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}