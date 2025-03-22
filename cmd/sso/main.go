package main

import (
	"jwt-go/internal/config"
	"jwt-go/pkg/logger"
)

func main() {
	log := logger.New()

	config, err := config.Parse()
	if err != nil {
		log.Error("config parsing error", logger.Error(err))
		return
	}

	_ = config
}
