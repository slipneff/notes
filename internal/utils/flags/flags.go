package flags

import (
	"errors"
	"flag"
)

const (
	configPathFlag = "config-path"
	envModeFlag    = "env-mode"
)

type CMDFlags struct {
	ConfigPath string
	EnvMode    string
}

func ParseFlags() (*CMDFlags, error) {
	configPath := flag.String(configPathFlag, "", "Configuration file path")
	envMode := flag.String(envModeFlag, "", "Environment mode")
	flag.Parse()

	if *configPath == "" {
		return nil, errors.New("Configuration file path was not found in application flags")
	}

	if *envMode == "" {
		return nil, errors.New("Environment mode was not found in application flags")
	}

	return &CMDFlags{ConfigPath: *configPath, EnvMode: *envMode}, nil
}

func MustParseFlags() *CMDFlags {
	flags, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	return flags
}
