package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SRVHost string
	SRVPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	CarsInfoURL string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	return &Config{
		SRVHost:     os.Getenv("SRV_HOST"),
		SRVPort:     os.Getenv("SRV_PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		CarsInfoURL: os.Getenv("CAR_INFO_API_URL"),
	}, nil
}
