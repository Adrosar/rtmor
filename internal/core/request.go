package core

import (
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// NewRequestHandler ...
func NewRequestHandler(tree *Tree, lm *LogMaster) *RequestHandler {
	return &RequestHandler{
		Tree: tree,
		LM:   lm,
	}
}

// RequestHandler ...
type RequestHandler struct {
	Tree *Tree
	LM   *LogMaster
}

// Handle ...
func (reqh RequestHandler) Handle(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	url := req.URL.String()
	rule := FindInTree(req.URL.Hostname(), url, reqh.Tree)

	if rule == nil {
		return req, nil
	}

	if rule.Mode >= 700 && rule.Mode < 800 {
		// 700-799 is reserved for -> ./internal/core/response.go
		return req, nil
	}

	if rule.ShowMatches {
		reqh.LM.Print('M', color.CyanString(`Rule: "`+rule.Name+`", URL -> `+url), "\n")
	}

	var res *http.Response

	switch rule.Mode {

	case RuleModeNotFound:
		res = NewRes404(req)

	case RuleModeRedirect:
		res = NewRes307(req, rule.Location)

	case RuleModeOK:
		if rule.Type == "text/javascript" {
			res = NewRes20X(req, AddLogToJS(rule.Body, rule.Name), rule.Type)
		} else {
			res = NewRes20X(req, rule.Body, rule.Type)
		}

	case RuleModeNoContent:
		res = NewRes20X(req, "", rule.Type)

	case RuleModeFile:
		text, _ := ReadTextFile(rule.Location)
		if rule.Type == "text/javascript" {
			res = NewRes20X(req, AddLogToJS(text, rule.Name), rule.Type)
		} else {
			res = NewRes20X(req, text, rule.Type)
		}

	}

	if res != nil {
		SetAntiCacheHeaders(res)
		SetInformationHeaders(res, rule)
		return req, res
	}

	return req, nil
}
