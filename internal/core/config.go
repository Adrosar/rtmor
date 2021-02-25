package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	// RuleModePASS ...
	RuleModePASS = 0

	// RuleModeOK ...
	RuleModeOK = 200

	// RuleModeNoContent ...
	RuleModeNoContent = 204

	// RuleModeFile ...
	RuleModeFile = 237

	// RuleModeRedirect ...
	RuleModeRedirect = 307

	// RuleModeNotFound ...
	RuleModeNotFound = 404

	// RuleModeNoCache ...
	RuleModeNoCache = 700
)

// Rule ...
type Rule struct {
	Name        string `yaml:"name"`
	Mode        int    `yaml:"mode"`
	ShowMatches bool   `yaml:"show_matches"`

	HostName string `yaml:"host_name"`
	RegExp   string `yaml:"reg_exp"`

	Location string `yaml:"location"`
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
		return nil, ExtendError("[sKIxye]", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, ExtendError("[TrVfXB]", err)
	}

	return &config, nil
}
