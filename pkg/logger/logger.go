package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	const op = "logger.New()"

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	return slog.New(logHandler)
}

func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
