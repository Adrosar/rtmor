package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// Proxy ...
type Proxy struct {
	Tree     *Tree
	ShowLogs bool
	Addr     string
	PHS      *goproxy.ProxyHttpServer
	Logger   *log.Logger
}

// NewProxy ...
func NewProxy() *Proxy {
	return &Proxy{
		Addr:     "127.0.0.1:8888",
		ShowLogs: false,
		Tree:     NewTree(),
		Logger:   log.New(os.Stdout, "", log.Ltime|log.Ldate),
		PHS:      goproxy.NewProxyHttpServer(),
	}
}

// InitProxy ...
func InitProxy(p *Proxy) {
	InitOutForLog(p.Logger, p.ShowLogs)
	p.PHS.Verbose = p.ShowLogs

	var mitmOrOK goproxy.FuncHttpsHandler = func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		hostnameAndPort := strings.Split(host, ":")
		if len(hostnameAndPort) > 0 {
			if IsHostNameExist(hostnameAndPort[0], p.Tree) {
				return goproxy.MitmConnect, host
			}
		}

		return goproxy.OkConnect, host
	}

	p.PHS.OnRequest().HandleConnect(mitmOrOK)
	p.PHS.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		url := req.URL.String()
		rule := FindInTree(req.URL.Hostname(), url, p.Tree)
		if rule == nil {
			return req, nil
		}

		if rule.ShowMatches {
			p.Logger.Println(color.CyanString(`Rule: "` + rule.Name + `", URL -> ` + url))
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
				res = NewRes20X(req, AddLogsToJS(rule.Body), rule.Type)
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
				res = NewRes20X(req, AddLogsToJS(text), rule.Type)
			} else {
				res = NewRes20X(req, text, rule.Type)
			}
			break
		}

		if res != nil {
			res.Header.Set("ETag", fmt.Sprint(UnixTime()))

			res.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
			res.Header.Set("Pragma", "no-cache")
			res.Header.Set("Expires", "0")

			res.Header.Set("Via", "RtMoR")
			res.Header.Set("X-Rtmor-Name", rule.Name)
			res.Header.Set("X-Rtmor-Mode", fmt.Sprint(rule.Mode))

			return req, res
		}

		return req, nil
	})
}

// RunProxy ...
func RunProxy(proxy *Proxy) error {
	err := http.ListenAndServe(proxy.Addr, proxy.PHS)
	if err != nil {
		return ExtendError("[4r8urC]", err)
	}

	return nil
}

// NewRes20X ...
func NewRes20X(req *http.Request, body string, contentType string) *http.Response {
	var res *http.Response

	if body == "" {
		res = goproxy.NewResponse(req, "text/plain", 204, "")
	} else {
		res = goproxy.NewResponse(req, contentType, 200, body)
	}

	return res
}

// NewRes404 ...
func NewRes404(req *http.Request) *http.Response {
	return goproxy.NewResponse(req, "text/plain", 404, "")
}

// NewRes307 ...
func NewRes307(req *http.Request, link string) *http.Response {
	res := goproxy.NewResponse(req, "text/plain", 307, "")
	res.Header.Set("Location", link)
	return res
}
