package core

import "regexp"

type Mode int

const (
	// RuleModePASS does almost nothing. Useful for testing rules.
	RuleModePASS Mode = 0

	// RuleModeOK returns the response set in "Rule.Body".
	// Uses the "Rule.Type" field for setting the MIME type.
	RuleModeOK Mode = 200

	// RuleModeNoContent simulates an empty response, code 204.
	// Uses the "Rule.Type" field for setting the MIME type.
	RuleModeNoContent Mode = 204

	// RuleModeFile returns a response from a file or other URL.
	// Uses the "Rule.Location" field as the source.
	// Uses the "Rule.Type" field for setting the MIME type.
	RuleModeFile Mode = 237

	// RuleModeRedirect forces a redirection.
	// Uses the "Rule.Location" field as a redirection target.
	RuleModeRedirect Mode = 307

	// RuleModeNotFound simulates a 404 error.
	RuleModeNotFound Mode = 404

	// RuleModeNoCache adds headers to the original response that prevent the browser from caching the resource.
	RuleModeNoCache Mode = 700
)

// Rule ...
type Rule struct {
	// Rule name.
	Name string `yaml:"name"`

	// Rule mode.
	Mode Mode `yaml:"mode"`

	// Show matching in terminal.
	ShowMatches bool `yaml:"show_matches"`

	// Hostname of the search query e.g.: "google.com"
	HostName string `yaml:"host_name"`

	// Regular expression for the query we want it to match.
	RegExp string `yaml:"reg_exp"`

	// Location for redirection or content.
	Location string `yaml:"location"`

	// Response Body.
	Body string `yaml:"body"`

	// MIME type for the response body, e.g.: "text/javascript"
	Type string `yaml:"type"`

	// Enable / disable the rule.
	Active bool `yaml:"active"`

	// Sets headers for response to prevent browser caching.
	PreventCache bool `yaml:"prevent_cache"`

	// Sets CORS for the response.
	CORS bool `yaml:"cors"`
}

// MatchTheRule ...
func (rule *Rule) MatchTheRule(text string) bool {
	re := regexp.MustCompile(rule.RegExp)
	return re.MatchString(text)
}
