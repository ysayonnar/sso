package server

import (
	"fmt"
	"jwt-go/api/handlers"
	"jwt-go/internal/config"
	"jwt-go/internal/database"
	"log/slog"
	"net/http"
	"time"
)

func New(cfg config.Server, logger *slog.Logger, storage *database.Storage) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HELLO NIGGA!")
	})

	mux.HandleFunc("/registration", handlers.Registration(logger, storage))

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      mux,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
}
