package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	PgHost     string
	PgPort     int
	PgUsername string
	PgPass     string
	PgDb       string
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
		PgHost:     PgHost,
		PgPort:     PgPortInt,
		PgUsername: PgUsername,
		PgPass:     PgPass,
		PgDb:       PgDb,
	}, nil
}
