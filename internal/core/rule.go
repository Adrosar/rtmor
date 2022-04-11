package core

import "regexp"

type Mode int

const (
	// RuleModePASS ...
	RuleModePASS Mode = 0

	// RuleModeOK ...
	RuleModeOK Mode = 200

	// RuleModeNoContent ...
	RuleModeNoContent Mode = 204

	// RuleModeFile ...
	RuleModeFile Mode = 237

	// RuleModeRedirect ...
	RuleModeRedirect Mode = 307

	// RuleModeNotFound ...
	RuleModeNotFound Mode = 404

	// RuleModeNoCache ...
	RuleModeNoCache Mode = 700
)

// Rule ...
type Rule struct {
	Name        string `yaml:"name"`
	Mode        Mode   `yaml:"mode"`
	ShowMatches bool   `yaml:"show_matches"`

	HostName string `yaml:"host_name"`
	RegExp   string `yaml:"reg_exp"`

	Location string `yaml:"location"`
	Body     string `yaml:"body"`
	Type     string `yaml:"type"`
	Active   bool   `yaml:"active"`
}

// MatchTheRule ...
func (rule *Rule) MatchTheRule(text string) bool {
	re := regexp.MustCompile(rule.RegExp)
	return re.MatchString(text)
}
