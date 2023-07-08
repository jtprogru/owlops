package logs

import (
	"fmt"
	"os"

	"github.com/jtprogru/owlops/internal/config"
	"golang.org/x/exp/slog"
)

type Logger struct {
	log *slog.Logger
}

func New() *Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}).
		WithAttrs([]slog.Attr{slog.String("version", config.Version)})
	logger := slog.New(jsonHandler)

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Log(msg string) {
	l.log.Info(fmt.Sprintf("%v\n", msg))
}
