package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

		var res *http.Response
		switch rule.Mode {

		case ShowURL:
			p.Logger.Println(color.CyanString("URL -> " + url))
			break

		case BlockRuleModed:
			res = NewRes404(req)
			res.Header.Set("x-rtmor-mode", fmt.Sprint(BlockRuleModed))
			return nil, res

		case RedirectRuleModed:
			res = NewRes307(req, rule.Location)
			res.Header.Set("x-rtmor-mode", fmt.Sprint(RedirectRuleModed))
			return nil, res

		case ContentRuleMode:
			if rule.Type == "text/javascript" {
				res = NewRes20X(req, AddLogsToJS(rule.Body), rule.Type)
			} else {
				res = NewRes20X(req, rule.Body, rule.Type)
			}

			res.Header.Set("x-rtmor-mode", fmt.Sprint(ContentRuleMode))
			return nil, res

		case FileRuleMode:
			text, _ := ReadTextFile(rule.Location)
			if rule.Type == "text/javascript" {
				res = NewRes20X(req, AddLogsToJS(text), rule.Type)
			} else {
				res = NewRes20X(req, text, rule.Type)
			}

			res.Header.Set("x-rtmor-mode", fmt.Sprint(FileRuleMode))
			return nil, res
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
	// https://stackoverflow.com/questions/33978216/create-http-response-instance-with-sample-body-string-in-golang
	res := &http.Response{
		Status:        "",
		StatusCode:    0,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header, 0),
	}

	if contentType == "" {
		res.Header.Set("Content-Type", "text/plain")
	} else {
		res.Header.Set("Content-Type", contentType)
	}

	if body == "" {
		res.StatusCode = 204
		res.Status = "204 No content"
	} else {
		res.StatusCode = 200
		res.Status = "200 OK"
	}

	return res
}

// NewRes404 ...
func NewRes404(req *http.Request) *http.Response {
	res := NewRes20X(req, "", "")
	res.StatusCode = 404
	res.Status = "404 Not Found"
	return res
}

// NewRes307 ...
func NewRes307(req *http.Request, link string) *http.Response {
	res := NewRes20X(req, "", "")
	res.StatusCode = 307
	res.Status = "307 Temporary Redirect"
	res.Header.Set("Location", link)
	return res
}
