package core

import (
	"net/http"
	"net/url"
	"strings"

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
	rule := reqh.Tree.FindURL(req.URL.Hostname(), req.URL.String())

	if rule == nil {
		return req, nil
	}

	if rule.Mode >= 700 && rule.Mode < 800 {
		// 700-799 is reserved for -> ./internal/core/response.go
		return req, nil
	}

	if rule.ShowMatches {
		reqh.LM.Print('M', color.CyanString(`Rule: "`+rule.Name+`", URL -> `+req.URL.String()), "\n")
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

	case RuleModeUrl:
		// The second part of the code for this rule is in file "response.go"
		ctx.UserData = rule

		location, err := url.Parse(rule.Location)
		if err != nil {
			return nil, NewRes404(req)
		}

		req.URL = location
		req.Host = location.Host

		return req, nil

	case RuleModeFile:
		var text string
		var err error

		protocol, path := SplitURL(rule.Location)
		if protocol == "file:" {
			text, err = ReadTextFile(path)
		} else if protocol == "http:" || protocol == "https:" {
			text, err = ReadTextFromURL(rule.Location)
		}

		if err != nil || len(text) == 0 {
			res = NewRes404(req)
		} else {
			if rule.Type == "text/javascript" {
				res = NewRes20X(req, AddLogToJS(text, rule.Name), rule.Type)
			} else {
				res = NewRes20X(req, text, rule.Type)
			}
		}
	}

	if res != nil {
		if rule.PreventCache {
			SetAntiCacheHeaders(res)
		}

		if rule.CORS {
			SetCORS(res, strings.TrimSpace(req.Header.Get("Origin")))
		}

		SetInformationHeaders(res, rule)
		return req, res
	}

	return req, nil
}
