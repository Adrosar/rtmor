package core

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// NewRequestHandler ...
func NewRequestHandler(tree *Tree, logger *log.Logger) *RequestHandler {
	return &RequestHandler{
		Tree:   tree,
		Logger: logger,
	}
}

// RequestHandler ...
type RequestHandler struct {
	Tree   *Tree
	Logger *log.Logger
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
		reqh.Logger.Println(color.CyanString(`Rule: "` + rule.Name + `", URL -> ` + url))
	}

	var res *http.Response

	switch rule.Mode {

	case RuleModeNotFound:
		res = NewRes404(req)
		break

	case RuleModeRedirect:
		res = NewRes307(req, rule.Location)
		break

	case RuleModeOK:
		if rule.Type == "text/javascript" {
			res = NewRes20X(req, AddLogToJS(rule.Body, rule.Name), rule.Type)
		} else {
			res = NewRes20X(req, rule.Body, rule.Type)
		}
		break

	case RuleModeNoContent:
		res = NewRes20X(req, "", rule.Type)
		break

	case RuleModeFile:
		text, _ := ReadTextFile(rule.Location)
		if rule.Type == "text/javascript" {
			res = NewRes20X(req, AddLogToJS(text, rule.Name), rule.Type)
		} else {
			res = NewRes20X(req, text, rule.Type)
		}
		break
	}

	if res != nil {
		SetAntiCacheHeaders(res)
		SetInformationHeaders(res, rule)
		return req, res
	}

	return req, nil
}
