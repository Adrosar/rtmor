package core

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	// ContentRuleMode ...
	ContentRuleMode = 211

	// LinkRuleMode ...
	LinkRuleMode = 212 // ToDo: Complete!

	// FileRuleMode ...
	FileRuleMode = 213 // ToDo: Complete!

	// RedirectRuleModed ...
	RedirectRuleModed = 307

	// BlockRuleModed ...
	BlockRuleModed = 404

	// ShowURL ...
	ShowURL = 600
)

// Rule ...
type Rule struct {
	Name     string `yaml:"name"`
	HostName string `yaml:"host_name"`
	RegExp   string `yaml:"reg_exp"`
	Location string `yaml:"location"`
	Mode     int    `yaml:"mode"`
	Body     string `yaml:"body"`
	Type     string `yaml:"type"`
	Active   bool   `yaml:"active"`
}

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
		return nil, errors.New(`[sKIxye] → ` + err.Error())
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.New(`[TrVfXB] → ` + err.Error())
	}

	return &config, nil
}
