package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Telegram `yaml:"telegram"`
}

type Telegram struct {
	Token string `yaml:"token" env-required:"true"`
}

// MustLoad loads config to a new Config instance and return it.
func MustLoad() *Config {
	_ = godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatalf("missed CONFIG_PATH parameter")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at: %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &config
}
