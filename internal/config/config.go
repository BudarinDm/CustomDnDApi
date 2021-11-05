package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Server   ServerConfig
		Postgres PostgresConfig
	}

	ServerConfig struct {
		Port string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		DBName   string
		Password string
		SSLMode  string
	}
)

func Init() (*Config, error) {
	var cfg Config

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg.Server.Port = viper.GetString("port")

	cfg.Postgres.Host = viper.GetString("db.host")
	cfg.Postgres.Port = viper.GetString("db.port")
	cfg.Postgres.Username = viper.GetString("db.username")
	cfg.Postgres.DBName = viper.GetString("db.dbname")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.SSLMode = viper.GetString("db.sslmode")

	return &cfg, nil
}
