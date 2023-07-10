package main

import (
	"fmt"

	"github.com/jtprogru/owlops/internal/config"
	"github.com/jtprogru/owlops/internal/logs"
)

func main() {
	// Get configuration
	cfg := config.GetConfig()

	// Create logger
	logger := logs.New(cfg.LogLevel)

	logger.Debug("app is running", "method", "cmd.owlops.main")
	logger.Info(fmt.Sprintf("%+v\n", cfg))
}
