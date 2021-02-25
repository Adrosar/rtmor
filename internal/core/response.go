package core

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// NewResponseHandler ...
func NewResponseHandler(tree *Tree, logger *log.Logger) *ResponseHandler {
	return &ResponseHandler{
		Tree:   tree,
		Logger: logger,
	}
}

// ResponseHandler ...
type ResponseHandler struct {
	Tree   *Tree
	Logger *log.Logger
}

// Handle ...
func (resh ResponseHandler) Handle(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	url := res.Request.URL.String()
	hostName := res.Request.URL.Hostname()
	rule := FindInTree(hostName, url, resh.Tree)
	if rule != nil {
		if rule.Mode == RuleModeNoCache {
			SetAntiCacheHeaders(res)
			SetInformationHeaders(res, rule)

			if rule.ShowMatches {
				resh.Logger.Println(color.YellowString(`Anti-buffering headers have been added to the response for the "` + url + `" address.`))
			}

			return res
		}
	}

	return res
}
