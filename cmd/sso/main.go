package main

import (
	"jwt-go/internal/config"
	"jwt-go/internal/database"
	"jwt-go/internal/logger"
)

func main() {
	log := logger.New()

	config, err := config.Parse()
	if err != nil {
		log.Error("config parsing error", logger.Error(err))
		return
	}
	log.Info("config parsed successfully")
	_ = config

	storage, err := database.Connect()
	if err != nil {
		log.Error("database conn error", logger.Error(err))
		return
	}

	_ = storage
}
