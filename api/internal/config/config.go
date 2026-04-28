package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/momoyo-droid/pismo/api/internal/utils"
)

// Config struct holds the configuration values for the application,
// including server port and database connection details.
type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// LoadConfig loads the configuration values from environment variables and validates them.
// It returns a Config struct and an error if any required configuration is missing or invalid.
func LoadConfig() (*Config, error) {
	env := ".env"
	if os.Getenv("APP_ENV") == "local" {
		env = ".env.local"
	}

	err := godotenv.Load(env)
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	if DBHost == "" {
		return nil, utils.ErrDBHostRequired
	}

	if DBUser == "" {
		return nil, utils.ErrDBUserRequired
	}

	if DBPassword == "" {
		return nil, utils.ErrDBPasswordRequired
	}

	if DBName == "" {
		return nil, utils.ErrDBNameRequired
	}

	if DBPort == "" {
		return nil, utils.ErrDBPortRequired
	}
	if _, err := strconv.Atoi(DBPort); err != nil {
		return nil, utils.ErrInvalidDBPort
	}

	if port == "" {
		return nil, utils.ErrPortRequired
	}

	if _, err := strconv.Atoi(port); err != nil {
		return nil, utils.ErrInvalidPort
	}

	cfg := &Config{
		Port:       port,
		DBHost:     DBHost,
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBName:     DBName,
		DBPort:     DBPort,
	}

	return cfg, nil
}
