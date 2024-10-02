package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string `env:"TELEGRAM_BOT_TOKEN" env-required:"true"`
	PgDsn    string `env:"PG_DSN" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	var cfg Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully")
	return &cfg, nil
}
