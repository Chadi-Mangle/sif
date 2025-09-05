package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type PostgresConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
	URL      string
}

type Config struct {
	DB     PostgresConfig
	Server ServerConfig
	Env    string
}

type ServerConfig struct {
	Host      string
	Port      string
	SecretKey string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Env: getEnvOrDefault("ENV", "development"),
		Server: ServerConfig{
			Host:      getEnvOrDefault("SERVER_HOST", "localhost"),
			Port:      getEnvOrDefault("HTTP_PORT", "3000"),
			SecretKey: getEnvOrDefault("SECRET_KEY", "default_secret_key"),
		},
		DB: PostgresConfig{
			Username: getEnvOrDefault("POSTGRES_USER", "postgres"),
			Password: getEnvOrDefault("POSTGRES_PASSWORD", "postgres"),
			Database: getEnvOrDefault("POSTGRES_DB", "go_web_startup"),
			Host:     getEnvOrDefault("POSTGRES_HOST", "localhost"),
			Port:     getEnvOrDefault("POSTGRES_PORT", "5432"),
			SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
		},
	}

	if dbURL := os.Getenv("DB_STRING"); dbURL != "" {
		cfg.DB.URL = dbURL
	} else if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		cfg.DB.URL = dbURL
	} else {
		cfg.DB.URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DB.Username,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Database,
			cfg.DB.SSLMode,
		)
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDatabaseURL() string {
	return c.DB.URL
}

func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}
