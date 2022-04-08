package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Rules []Rule `yaml:"rules"`
}

// ReadConfig ...
func ReadConfig(pathToFile string) (*Config, error) {
	var config Config
	var err error

	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nil, ExtendError("[sKIxye]", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, ExtendError("[TrVfXB]", err)
	}

	return &config, nil
}
