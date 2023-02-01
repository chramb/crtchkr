package util

import (
	"github.com/BurntSushi/toml"
	"os"
)

type ConfigStruct struct {
	Certs   []string
	Discord map[string]struct {
		Url     string
		Request string
	}
	Mail map[string]struct {
		Server  string
		From    string
		To      []string
		Message string
		Auth    struct {
			Identity string
			Username string
			Password string
			Host     string
		}
	}
	Exec map[string]struct {
		Script string
	}
}

func GetConfig(path string) (ConfigStruct, toml.MetaData, error) {
	var config ConfigStruct
	configFile := path
	if _, err := os.Stat(configFile); err != nil {
		return ConfigStruct{}, toml.MetaData{}, err
	}

	meta, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return ConfigStruct{}, toml.MetaData{}, err
	}
	return config, meta, nil
}
