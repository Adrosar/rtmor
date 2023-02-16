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
		resh.LM.Print('B', color.RedString("[q4sI33]"), color.YellowString("Response is `nil`"))
		return nil
	}

	if res.Request == nil {
		resh.LM.Print('B', color.RedString("[Wm79he]"), color.YellowString("Request is `nil`"))
		return nil
	}

	if res.Request == nil {
		resh.LM.Print('B', color.RedString("[u2q5cv]"), color.YellowString("URL is `nil`"))
		return nil
	}

	var reqUrl string
	var reqHostName string
	var rule *Rule

	if ctx.UserData == nil {
		reqUrl = res.Request.URL.String()
		reqHostName = res.Request.URL.Hostname()
		rule = resh.Tree.FindURL(reqHostName, reqUrl)
	} else {
		rule = ctx.UserData.(*Rule)
	}

	if rule != nil {
		if rule.Mode == RuleModeUrl {
			// The first part of the code for this rule is in file "request.go"

			if rule.PreventCache {
				SetAntiCacheHeaders(res)
			}

			if rule.CORS {
				SetCORS(res, "")
			}

			if len(rule.Type) > 0 {
				res.Header.Set("Content-Type", rule.Type)
			}

			SetInformationHeaders(res, rule)

			return res
		}

		if rule.Mode == RuleModeNoCache {
			SetAntiCacheHeaders(res)
			SetInformationHeaders(res, rule)

			if rule.ShowMatches {
				resh.LM.Print('M', color.YellowString(`Anti-buffering headers have been added to the response for the "`+reqUrl+`" address.`), "\n")
			}

			return res
		}
	}

	return res
}
