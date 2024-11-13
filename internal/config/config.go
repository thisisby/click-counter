package config

import (
	"click-counter/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env string `env:"env"`

	Host string `yaml:"host"`
	Port string `yaml:"port"`

	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

var Config config

func (c *config) MustInitializeConfig() {
	err := cleanenv.ReadConfig("./internal/config/config.yml", &Config)
	if err != nil {
		logger.ZeroLogger.Fatal().Msgf("config -> MustInitializeConfig -> cleanenv.ReadConfig: %v", err)
	}

}
