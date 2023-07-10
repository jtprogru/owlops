package logs

import (
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

const (
	op = "internal.logs.New"
)

func New(logLevel string) *slog.Logger {
	var logLvl slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		logLvl = slog.LevelDebug
	case "error":
		logLvl = slog.LevelError
	default:
		logLvl = slog.LevelInfo
	}

	jsonHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: logLvl,
		})

	logger := slog.New(jsonHandler)
	logger.Debug("logger initialized", "method", op)

	return logger
}
