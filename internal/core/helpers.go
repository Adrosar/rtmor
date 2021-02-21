package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

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

// AddLogsToJS adds "console.log()" with the "[RtMoR]" message to the beginning of the JavaScript code.
func AddLogsToJS(jsCode string) string {
	return fmt.Sprint(`if (typeof console != "undefined") { console.log("[RtMoR]") };`, "\n", jsCode)
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
