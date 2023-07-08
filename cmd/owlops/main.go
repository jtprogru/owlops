package main

import (
	"fmt"

	"github.com/jtprogru/owlops/internal/config"
	"github.com/jtprogru/owlops/internal/logs"
)

func main() {

	cfg := config.GetConfig()
	logger := logs.New(cfg.LogLevel)
	logger.Info("config initialized")
	logger.Info(fmt.Sprintf("%+v\n", cfg))
}
