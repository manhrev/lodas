package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	configFile = "./../config/config.yaml"
	configType = "yml"
)

type localConfig struct {
}

func (c *localConfig) New() (*Configuration, error) {
	viper.SetConfigType(configType)
	configPath, err := filepath.Abs(configFile)

	if err != nil {
		return nil, err
	}
	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &Configuration{}

	err = viper.Unmarshal(conf)

	if err != nil {
		return nil, err
	}
	return conf, nil
}
