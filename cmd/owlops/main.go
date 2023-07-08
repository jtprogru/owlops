package main

import (
	"fmt"

	"github.com/jtprogru/owlops/internal/config"
	"github.com/jtprogru/owlops/internal/logs"
)

func main() {
	logger := logs.New()

	cfg := config.GetConfig()
	logger.Log("config initialized")
	logger.Log(fmt.Sprintf("%v\n", cfg))
}
