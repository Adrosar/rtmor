package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rtmor/internal/core"

	"github.com/fatih/color"
)

func main() {
	var isStart bool
	flag.BoolVar(&isStart, "start", false, "Start the proxy server")

	opt := core.NewOptions()
	flag.BoolVar(&opt.ShowLogs, "log", false, "Show logs")
	flag.StringVar(&opt.ConfigFile, "cfg", "", "Path to the configuration file (YAML)")
	flag.StringVar(&opt.ProxyServerAddr, "listen", "127.0.0.1:8888", "The address on which the proxy server should listen")

	var toRepo bool
	flag.BoolVar(&toRepo, "repo", false, "Open the repository website")

	var toHelp bool
	flag.BoolVar(&toHelp, "help", false, "Display help")

	flag.Parse()

	if toHelp {
		fmt.Println("")
		fmt.Println("Real-time Modification of Requests")
		fmt.Println(" • version: 0.3.0 (2021-02-24-2236)")
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

	logger := log.New(os.Stdout, "", log.Ltime|log.Ldate)
	core.InitOutForLog(logger, opt.ShowLogs)

	if toRepo {
		err3 := core.OpenURLinBrowser("https://github.com/Adrosar/rtmor")
		if err3 != nil {
			logger.Fatalln(color.RedString(fmt.Sprint("[cdEQa9] →", err3)))
		}

		os.Exit(0)
	}

	if isStart == false {
		fmt.Println("Run the `rtmor -help` to see help")
		os.Exit(0)
	}

	proxy := core.NewProxy()
	proxy.ShowLogs = opt.ShowLogs
	proxy.Addr = opt.ProxyServerAddr

	if opt.ConfigFile != "" {
		conf, err1 := core.ReadConfig(opt.ConfigFile)
		if err1 != nil {
			logger.Fatalln(color.RedString(fmt.Sprint("[Uqd3CI] →", err1)))
		}

		for _, rule := range conf.Rules {
			ok := core.AddToTree(rule, proxy.Tree)
			if ok {
				logger.Println(`Rule "` + rule.Name + `" has been loaded ` + color.GreenString(`:)`))
			} else {
				logger.Println(`"` + rule.Name + `" rule failed to load ` + color.RedString(`:(`))
			}
		}
	} else {
		logger.Println(color.YellowString("Configuration file not selected!"))
	}

	core.InitProxy(proxy)
	logger.Println("The proxy server is listening at address → " + opt.ProxyServerAddr)
	err2 := core.RunProxy(proxy)
	if err2 != nil {
		logger.Fatalln(color.RedString(fmt.Sprint("[WxC8Y7] →", err2)))
	}
}
