package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
	UserHost          string
	UserPort          string
	AccommodationHost string
	AccommodationPort string
	ProfileHost       string
	ProfilePort       string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:              os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		ProfileHost:       os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:       os.Getenv("PROFILE_SERVICE_PORT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}
