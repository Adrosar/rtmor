package core

import (
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

// ProxyCore ...
type ProxyCore struct {
	Tree     *Tree
	ShowLogs bool
	Addr     string
	PHS      *goproxy.ProxyHttpServer
	LM       *LogMaster
}

// NewProxyCore ...
func NewProxyCore(logMaster *LogMaster) *ProxyCore {
	return &ProxyCore{
		Addr:     "127.0.0.1:8888",
		ShowLogs: false,
		Tree:     NewTree(),
		LM:       logMaster,
		PHS:      goproxy.NewProxyHttpServer(),
	}
}

// Init ...
func (pc *ProxyCore) Init() {
	if pc.LM.Map['V'] || pc.LM.ShowAll {
		pc.PHS.Verbose = true
		InitOutForLog(pc.PHS.Logger, true)
	} else if pc.LM.Map['W'] {
		pc.PHS.Verbose = false
		InitOutForLog(pc.PHS.Logger, true)
	} else {
		pc.PHS.Verbose = false
		InitOutForLog(pc.PHS.Logger, false)
	}

	pc.PHS.Tr.MaxConnsPerHost = 5
	pc.PHS.Tr.MaxIdleConns = 1000
	pc.PHS.Tr.MaxIdleConnsPerHost = 5
	pc.PHS.Tr.IdleConnTimeout = time.Second * 120
	pc.PHS.Tr.ResponseHeaderTimeout = time.Second * 30

	pc.PHS.OnResponse().Do(NewResponseHandler(pc.Tree, pc.LM))
	pc.PHS.OnRequest().Do(NewRequestHandler(pc.Tree, pc.LM))
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
