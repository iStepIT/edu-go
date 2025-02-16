package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}
type DbConfig struct {
	Dsn string
}
type AuthConfig struct {
	Secret string
}

type VerifyMailConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}

func LoadVerifyConfig() *VerifyMailConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &VerifyMailConfig{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
