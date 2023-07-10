package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"`
	Port     uint16 `yaml:"port" env:"PORT"`
}

const (
	op                 = "internal.config.New"
	configDefaulPath   = "config/config.local.yaml"
	EnvConfigPathName  = "CONFIG_PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func init() {
	flag.StringVar(&configPath, FlagConfigPathName, configDefaulPath, "this is app config file")
	flag.Parse()
}

func GetConfig() *Config {
	instance = &Config{}

	once.Do(func() {

		log.Println(fmt.Sprintf("%s config init", op))

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Println(fmt.Sprintf("%s config path is required", op))
			return
		}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "OwlOps - Duty Service"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Println(help)
			log.Println(fmt.Sprintf("%s err: %s", op, err.Error()))
			return
		}
	})

	return instance
}
