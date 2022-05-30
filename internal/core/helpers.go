package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/fatih/color"
)

// OpenURLinBrowser opens URL in default web browser.
// I used the code from https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func OpenURLinBrowser(url string) error {
	var err error = nil

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errors.New(`[Uv9ELW] → The "` + color.CyanString(url) + `" address could not be opened. Unsupported platform!`)
	}

	return err
}

// AddLogToJS ...
func AddLogToJS(jsCode string, ruleName string) string {
	return fmt.Sprint(`if (typeof console != "undefined") { console.log("[RtMoR] Rule: `+ruleName+`") };`, "\n", jsCode)
}

// InitOutForLog ...
func InitOutForLog(logger interface{}, show bool) {
	lgr, ok := logger.(*log.Logger)
	if !ok {
		return
	}

	if show {
		if runtime.GOOS == "windows" {
			lgr.SetOutput(color.Output)
		} else {
			lgr.SetOutput(os.Stdout)
		}
	} else {
		lgr.SetOutput(ioutil.Discard)
	}
}

// ReadTextFile ...
func ReadTextFile(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", ExtendError("[HFx9Mq]", err)
	}

	return string(data), nil
}

// ExtendError ...
func ExtendError(code string, err error) error {
	return errors.New(code + ` → ` + err.Error())
}

// UnixTime ...
func UnixTime() int64 {
	now := time.Now()
	return now.Unix()
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

// SetAntiCacheHeaders ...
func SetAntiCacheHeaders(res *http.Response) {
	res.Header.Set("ETag", fmt.Sprint(UnixTime()))
	res.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	res.Header.Set("Pragma", "no-cache")
	res.Header.Set("Expires", "0")
}

// SetCORS ...
func SetCORS(res *http.Response, origin string) {
	if len(origin) > 0 {
		res.Header.Set("Access-Control-Allow-Origin", origin)
		res.Header.Set("Access-Control-Allow-Credentials", "true")
	} else {
		res.Header.Set("Access-Control-Allow-Origin", "*")
	}

	res.Header.Set("Access-Control-Allow-Methods", "HEAD, OPTIONS, GET, POST")
	res.Header.Set("Access-Control-Allow-Headers", "Content-Type")
}

// SetInformationHeaders ...
func SetInformationHeaders(res *http.Response, rule *Rule) {
	res.Header.Set("Via", "RtMoR")
	res.Header.Set("X-Rtmor-Name", rule.Name)
	res.Header.Set("X-Rtmor-Mode", fmt.Sprint(rule.Mode))
}

// ReadAsString ...
func ReadAsString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()
}

// SplitURL splits the URL into a protocol and the rest
//
// EXAMPLEs:
//   [1] "http://127.0.0.1:8080" => "http:", "//127.0.0.1:8080"
//   [2] "https://google.com" => "https:", "//google.com"
//   [3] "file:configs/sample.yaml" => "file:", "configs/sample.yaml"
//
func SplitURL(location string) (protocol string, rest string) {
	prefix := location[0:5]
	if prefix == "file:" || prefix == "http:" {
		return prefix, location[5:]
	} else {
		prefix := location[0:6]
		if prefix == "https:" {
			return prefix, location[6:]
		}
	}

	return "", location
}

// ReadTextFromURL ...
func ReadTextFromURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", ExtendError("[gwfZLQ]", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", ExtendError("[hD9qQn]", err)
	}

	return string(body), nil
}
