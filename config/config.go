package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrMissingKey = errors.New("missing key")
)

type Config struct {
	DatabaseURL string
	ServerAddr string
}

func InitConfig() error {
	return godotenv.Load()
}

func LoadConfig() (*Config, error) {
	ret := Config{}
	
	if (os.Getenv("DATABASE_URL") == "") {
		return nil, ErrMissingKey
	}
	ret.DatabaseURL = os.Getenv("DATABASE_URL")

	if (os.Getenv("SERVER_ADDR") == "") {
		return nil, ErrMissingKey
	}
	ret.ServerAddr = os.Getenv("SERVER_ADDR")

	return &ret, nil
}