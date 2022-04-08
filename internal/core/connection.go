package core

import (
	"strings"

	"github.com/elazarl/goproxy"
)

// NewConnectionHandler ...
func NewConnectionHandler(tree *Tree) *ConnectionHandler {
	return &ConnectionHandler{
		Tree: tree,
	}
}

// ConnectionHandler ...
type ConnectionHandler struct {
	Tree *Tree
}

// HandleConnect ...
func (conh ConnectionHandler) HandleConnect(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	hostnameAndPort := strings.Split(host, ":")
	if len(hostnameAndPort) > 0 {
		if conh.Tree.IsHostNameExist(hostnameAndPort[0]) {
			return goproxy.MitmConnect, host
		}
	}

	return goproxy.OkConnect, host
}
