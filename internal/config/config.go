package config

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"golang.org/x/exp/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port uint16 `yaml:"port" env:"PORT"`
}

const (
	// configPath         = "config/config.yaml"
	op                 = "internal.config.New"
	Version            = "v0.1.0"
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "config/config.local.yaml", "this is app config file")
		flag.Parse()

		slog.Info(fmt.Sprintf("%s config init", op))

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			slog.Info(fmt.Sprintf("%s config path is required", op))
			return
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "The Art of Development - Production Service"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			slog.Info(help)
			slog.Info(fmt.Sprintf("%s err: %s", op, err.Error()))
			return
		}
	})
	return instance
}
