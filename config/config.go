package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	ErrMissingKey = errors.New("missing key")
)

type Config struct {
	DatabaseURL string
	ServerAddr  string
	CorsDomain  string

	JWTIssuer          string
	SellerAccessSecret string
	UserAccessSecret   string
	AdminAccessSecret  string

	AccessTokenLifespan time.Duration
}

func InitConfig() error {
	return godotenv.Load()
}

func LoadConfig() (*Config, error) {
	ret := Config{}

	if os.Getenv("DATABASE_URL") == "" {
		return nil, ErrMissingKey
	}
	ret.DatabaseURL = os.Getenv("DATABASE_URL")

	if os.Getenv("SERVER_ADDR") == "" {
		return nil, ErrMissingKey
	}
	ret.ServerAddr = os.Getenv("SERVER_ADDR")

	if os.Getenv("CORS_DOMAIN") == "" {
		return nil, ErrMissingKey
	}
	ret.CorsDomain = os.Getenv("CORS_DOMAIN")

	if os.Getenv("JWT_ISSUER") == "" {
		return nil, ErrMissingKey
	}
	ret.JWTIssuer = os.Getenv("JWT_ISSUER")

	if os.Getenv("SELLER_ACCESS_SECRET") == "" {
		return nil, ErrMissingKey
	}
	ret.SellerAccessSecret = os.Getenv("SELLER_ACCESS_SECRET")

	if os.Getenv("USER_ACCESS_SECRET") == "" {
		return nil, ErrMissingKey
	}
	ret.UserAccessSecret = os.Getenv("USER_ACCESS_SECRET")

	if os.Getenv("ADMIN_ACCESS_SECRET") == "" {
		return nil, ErrMissingKey
	}
	ret.AdminAccessSecret = os.Getenv("ADMIN_ACCESS_SECRET")

	s := os.Getenv("ACCESS_TOKEN_LIFESPAN")
	i, err := strconv.Atoi(s)
	if err != nil {
		return &ret, err
	}
	ret.AccessTokenLifespan = time.Duration(i) * time.Minute

	return &ret, nil
}
