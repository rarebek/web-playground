package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	PG_HOST     string
	PG_PORT     int
	PG_USERNAME string
	PG_PASS     string
	PG_DB       string
}

func Load() (*Config, error) {
	PgHost := os.Getenv("PG_HOST")
	PgPort := os.Getenv("PG_PORT")
	PgUsername := os.Getenv("PG_USERNAME")
	PgPass := os.Getenv("PG_PASS")
	PgDb := os.Getenv("PG_DB")

	PgPortInt, err := strconv.Atoi(PgPort)
	if err != nil {
		return nil, errors.New("PG PORT must be consists of only numbers")
	}

	return &Config{
		PG_HOST:     PgHost,
		PG_PORT:     PgPortInt,
		PG_USERNAME: PgUsername,
		PG_PASS:     PgPass,
		PG_DB:       PgDb,
	}, nil
}
