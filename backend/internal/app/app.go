package app

import (
	"context"
	"erp/internal/config"
	"erp/internal/database"
	"erp/internal/router"
	"log"
)

func Run() {
	cfg := config.Load()
	db, err := database.ConnectMySQL(cfg.Database)
	if err != nil {
		log.Fatalf("connect mysql: %v", err)
	}
	if _, err := database.ConnectRedis(context.Background(), cfg.Redis); err != nil {
		log.Printf("redis unavailable: %v", err)
	}
	if cfg.Database.AutoMigrate {
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("auto migrate: %v", err)
		}
		if err := database.Seed(db); err != nil {
			log.Fatalf("seed data: %v", err)
		}
	}
	r := router.New(db, cfg)
	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
