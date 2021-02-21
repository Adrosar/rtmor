package core

import "regexp"

// MatchTheRule ...
func MatchTheRule(rule *Rule, text string) bool {
	re := regexp.MustCompile(rule.RegExp)
	return re.MatchString(text)
}
