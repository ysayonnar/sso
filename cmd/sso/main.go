package main

import (
	"fmt"
	"jwt-go/api/server"
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

	storage, err := database.Connect()
	if err != nil {
		log.Error("database conn error", logger.Error(err))
		return
	}
	log.Info("database connected successfully")

	s := server.New(config.Server, log, storage)

	log.Info(fmt.Sprintf("server started on %s:%d", config.Server.Host, config.Server.Port))
	err = s.ListenAndServe()
	if err != nil {
		log.Error("server rising error", logger.Error(err))
		return
	}
}
