package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	JWT      JWTConfig      `yaml:"jwt"`
}

type AppConfig struct {
	Env  string `yaml:"env"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	DSN         string `yaml:"dsn"`
	AutoMigrate bool   `yaml:"autoMigrate"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret         string        `yaml:"secret"`
	ExpiresInHours int           `yaml:"expiresInHours"`
	ExpiresIn      time.Duration `yaml:"-"`
}

func Load() Config {
	cfg := defaultConfig()
	path := os.Getenv("ERP_CONFIG")
	if path == "" {
		path = "config.yml"
	}
	content, err := os.ReadFile(path)
	if err == nil {
		if err := yaml.Unmarshal(content, &cfg); err != nil {
			panic(err)
		}
	}
	cfg.JWT.ExpiresIn = time.Duration(cfg.JWT.ExpiresInHours) * time.Hour
	return cfg
}

func defaultConfig() Config {
	return Config{
		App: AppConfig{
			Env:  "development",
			Port: "8080",
		},
		Database: DatabaseConfig{
			DSN:         "erp:erp@tcp(127.0.0.1:3306)/erp?charset=utf8mb4&parseTime=True&loc=Local",
			AutoMigrate: true,
		},
		Redis: RedisConfig{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		},
		JWT: JWTConfig{
			Secret:         "change-this-development-secret",
			ExpiresInHours: 24,
			ExpiresIn:      24 * time.Hour,
		},
	}
}
