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
func InitOutForLog(logger *log.Logger, show bool) {
	if show {
		if runtime.GOOS == "windows" {
			logger.SetOutput(color.Output)
		} else {
			logger.SetOutput(os.Stdout)
		}
	} else {
		logger.SetOutput(ioutil.Discard)
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
	return `alert(123); ` + buf.String()
}
