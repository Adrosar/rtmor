package core

import (
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// NewResponseHandler ...
func NewResponseHandler(tree *Tree, lm *LogMaster) *ResponseHandler {
	return &ResponseHandler{
		Tree: tree,
		LM:   lm,
	}
}

// ResponseHandler ...
type ResponseHandler struct {
	Tree *Tree
	LM   *LogMaster
}

// Handle ...
func (resh ResponseHandler) Handle(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if res == nil {
		resh.LM.Print('B', color.RedString("[q4sI33]"))
		return nil
	}

	if res.Request != nil && res.Request.URL != nil {
		url := res.Request.URL.String()
		hostName := res.Request.URL.Hostname()
		rule := FindInTree(hostName, url, resh.Tree)

		if rule != nil {
			if rule.Mode == RuleModeNoCache {
				SetAntiCacheHeaders(res)
				SetInformationHeaders(res, rule)

				if rule.ShowMatches {
					resh.LM.Print('M', color.YellowString(`Anti-buffering headers have been added to the response for the "`+url+`" address.`), "\n")
				}

				return res
			}
		}
	} else {
		resh.LM.Print('B', color.YellowString("[fAqg6N]"))
	}

	return res
}
