package core

// Tree ...
type Tree struct {
	Data map[string][]Rule
}

// NewTree ...
func NewTree() *Tree {
	return &Tree{
		Data: map[string][]Rule{},
	}
}

// AddToTree added "rule" to "tree".
func AddToTree(rule Rule, tree *Tree) (ok bool) {
	l := len(rule.HostName)
	if l == 0 {
		return false
	}

	tree.Data[rule.HostName] = append(tree.Data[rule.HostName], rule)
	return true
}

// FindInTree ...
func FindInTree(hostName string, text string, tree *Tree) *Rule {
	list := tree.Data[hostName]
	if list == nil {
		return nil
	}

	for _, rule := range list {
		if rule.Active == false {
			continue
		}

		if MatchTheRule(&rule, text) {
			return &rule
		}
	}

	return nil
}

// IsHostNameExist ...
func IsHostNameExist(hostName string, tree *Tree) bool {
	if hostName == "" {
		return false
	}

	list := tree.Data[hostName]
	if list == nil {
		return false
	}

	return true
}
