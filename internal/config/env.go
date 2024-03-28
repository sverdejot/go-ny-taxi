package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}

// TODO: need proper init
func Init() (Config, error) {
	config := Config{}

	config.Database.Host = os.Getenv("NYTAXI_DB_HOST")
	config.Database.Port, _ = strconv.Atoi(os.Getenv("NYTAXI_DB_PORT"))
	config.Database.User = os.Getenv("NYTAXI_DB_USER")
	config.Database.Password = os.Getenv("NYTAXI_DB_PASS")
	config.Database.Database = os.Getenv("NYTAXI_DB_NAME")

	return config, nil
}
