package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	ENV_MODE_DEVELOPMENT = iota + 1
	ENV_MODE_PRODUCTION  = iota + 1
	ENV_MODE_STAGE       = iota + 1
)

const (
	envModeDevelopmentStr = "development"
	envModeProductionStr  = "production"
	envModeStageStr       = "stage"
)

type Config struct {
	EnvMode   uint8
	Host      string
	Port      uint16
	DB        Database
}

type Database struct {
	Host     string
	Port     uint16
	Username string
	Name     string
	Password string
	Sslmode  string
}

func LoadConfig(envMode, path string) (*Config, error) {
	mode, err := validateEnvMode(envMode)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	config.EnvMode = mode

	return config, nil
}

func MustLoadConfig(envMode, path string) *Config {
	config, err := LoadConfig(envMode, path)
	if err != nil {
		panic(err)
	}

	return config
}

func validateEnvMode(envMode string) (uint8, error) {
	var mode uint8
	switch envMode {
	case envModeDevelopmentStr:
		mode = ENV_MODE_DEVELOPMENT
	case envModeProductionStr:
		mode = ENV_MODE_PRODUCTION
	case envModeStageStr:
		mode = ENV_MODE_STAGE
	default:
		return mode, errors.New("Unknown environment mode")
	}

	return mode, nil
}
