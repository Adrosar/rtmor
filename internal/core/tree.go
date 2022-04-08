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

// AddRule added "rule" to "tree".
func (tree *Tree) AddRule(rule Rule) (ok bool) {
	l := len(rule.HostName)
	if l == 0 {
		return false
	}

	tree.Data[rule.HostName] = append(tree.Data[rule.HostName], rule)
	return true
}

// FindURL ...
func (tree *Tree) FindURL(hostName string, text string) *Rule {
	list := tree.Data[hostName]
	if list == nil {
		return nil
	}

	for _, rule := range list {
		if !rule.Active {
			continue
		}

		if rule.MatchTheRule(text) {
			return &rule
		}
	}

	return nil
}

// IsHostNameExist ...
func (tree *Tree) IsHostNameExist(hostName string) bool {
	if hostName == "" {
		return false
	}

	return tree.Data[hostName] != nil
}
