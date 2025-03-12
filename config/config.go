package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB           *DBConfig     `json:"db"`
	Server       *ServerConfig `json:"server"`
	JWTSecretKey string
}

type DBConfig struct {
	Driver        string `json:"driver,omitempty"`
	DSN           string `json:"dsn,omitempty"`
	MigrationsDir string `json:"migration_dir,omitempty"`
}

type ServerConfig struct {
	IPAddress string `json:"ipaddress"`
	Port      string `json:"port"`
}

type JWTSecretKey string

var config Config

func init() {
	// Load config from file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConf := &DBConfig{
		Driver:        os.Getenv("DB_DRIVER"),
		DSN:           os.Getenv("DB_DSN"),
		MigrationsDir: os.Getenv("DB_MIGRATION_DIR"),
	}

	serverConf := &ServerConfig{
		IPAddress: os.Getenv("HOST"),
		Port:      os.Getenv("PORT"),
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	config = Config{
		DB:           dbConf,
		Server:       serverConf,
		JWTSecretKey: JWTSecretKey,
	}
}

func GetConfig() *Config {
	return &config
}
