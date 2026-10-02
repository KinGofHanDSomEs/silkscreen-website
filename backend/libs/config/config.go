package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App AppConfig `env-prefix:"APP_"`
	DB  DBConfig  `env-prefix:"DB_"`
}

type AppConfig struct {
	Port int `env:"APP_PORT" envDefault:"8080"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST" env-default:"db"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	Database string `env:"DB_DATABASE" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

func MustLoad(path string) *Config {
	var cfg Config

	if path == "" {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("invalid read env: %v", err)
		}
	} else if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("invalid read config: %v", err)
	}

	help, err := cleanenv.GetDescription(&cfg, nil)
	if err != nil {
		log.Fatalf("invalid description env: %v", err)
	}

	log.Println(help)

	return &cfg
}
