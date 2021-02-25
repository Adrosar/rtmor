package core

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

// ProxyCore ...
type ProxyCore struct {
	Tree     *Tree
	ShowLogs bool
	Addr     string
	PHS      *goproxy.ProxyHttpServer
	Logger   *log.Logger
}

// NewProxyCore ...
func NewProxyCore() *ProxyCore {
	return &ProxyCore{
		Addr:     "127.0.0.1:8888",
		ShowLogs: false,
		Tree:     NewTree(),
		Logger:   log.New(os.Stdout, "", log.Ltime|log.Ldate),
		PHS:      goproxy.NewProxyHttpServer(),
	}
}

// Init ...
func (pc *ProxyCore) Init() {
	InitOutForLog(pc.Logger, pc.ShowLogs)
	pc.PHS.Verbose = pc.ShowLogs

	pc.PHS.OnResponse().Do(NewResponseHandler(pc.Tree, pc.Logger))
	pc.PHS.OnRequest().Do(NewRequestHandler(pc.Tree, pc.Logger))
	pc.PHS.OnRequest().HandleConnect(NewConnectionHandler(pc.Tree))
}

// Run ...
func (pc *ProxyCore) Run() error {
	err := http.ListenAndServe(pc.Addr, pc.PHS)
	if err != nil {
		return ExtendError("[4r8urC]", err)
	}

	return nil
}
