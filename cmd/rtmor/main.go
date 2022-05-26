package main

import (
	"flag"
	"fmt"
	"os"
	"rtmor/internal/core"

	"github.com/fatih/color"
)

const ver = "0.8.0 (2022-05-26-2333)"

const logDesc = `Shows the logs. Use:
'-log B' → Basic logs.
'-log M' → Rules matching logs.
'-log W' → Warnings and Errors from https://github.com/elazarl/goproxy
'-log V' → Verbose log from https://github.com/elazarl/goproxy
           (implies the use of the 'W' flag)
'-log BMV' → Show all logs.
'-log BMW' → Most common use.`

const cfgDesc = `Path to the configuration file (YAML)
See example: https://github.com/Adrosar/rtmor/blob/0.7.3/configs/sample.yam`

const listenDesc = `The address on which the proxy server should listen.
To listen on all interfaces (network adapters), use '-listen 0.0.0.0:8888'
`

func main() {
	var isStart bool
	flag.BoolVar(&isStart, "start", false, "Start the proxy server")

	var logFlags string
	flag.StringVar(&logFlags, "log", "", logDesc)

	var configFileName string
	flag.StringVar(&configFileName, "cfg", "", cfgDesc)

	var proxyServerAddr string
	flag.StringVar(&proxyServerAddr, "listen", "127.0.0.1:8888", listenDesc)

	var toRepo bool
	flag.BoolVar(&toRepo, "repo", false, "Open the repository website")

	var toHelp bool
	flag.BoolVar(&toHelp, "help", false, "Display help")

	flag.Parse()

	if toHelp {
		fmt.Println("")
		fmt.Println("Real-time Modification of Requests")
		fmt.Println(" • version:", ver)
		fmt.Println(" • author: Adrian Gargula")
		fmt.Println(" • Go-compatible BSD license")
		fmt.Println("")
		fmt.Println("Documentation of use:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Minimal command example:")
		fmt.Println(" → rtmor -start")
		fmt.Println("")
		os.Exit(0)
	}

	lm := core.NewLogMaster()
	lm.Parse(logFlags)

	if toRepo {
		err3 := core.OpenURLinBrowser("https://github.com/Adrosar/rtmor")
		if err3 != nil {
			lm.Print('B', color.RedString(fmt.Sprint("[cdEQa9] →", err3)), "\n")
			os.Exit(1)
		}

		os.Exit(0)
	}

	if !isStart {
		fmt.Println("Run the `rtmor -help` to see help")
		os.Exit(0)
	}

	lm.Print('B', color.HiCyanString("RtMoR "+ver), "\n")

	pc := core.NewProxyCore(lm)
	pc.Addr = proxyServerAddr

	if configFileName != "" {
		conf, err1 := core.ReadConfig(configFileName)
		if err1 != nil {
			lm.Print('B', color.RedString(fmt.Sprint("[Uqd3CI] →", err1)), "\n")
			os.Exit(2)
		}

		for _, rule := range conf.Rules {
			ok := pc.Tree.AddRule(rule)
			if ok {
				lm.Print('B', `Rule "`+rule.Name+`" has been loaded `+color.GreenString(`:)`), "\n")
			} else {
				lm.Print('B', `"`+rule.Name+`" rule failed to load `+color.RedString(`:(`), "\n")
			}
		}
	} else {
		lm.Print('B', color.YellowString("Configuration file not selected!"), "\n")
	}

	pc.Init()
	lm.Print('B', "The proxy server is listening at address → "+proxyServerAddr, "\n")
	err2 := pc.Run()
	if err2 != nil {
		lm.Print('B', color.RedString(fmt.Sprint("[WxC8Y7] → ", err2)), "\n")
		os.Exit(3)
	}
}
