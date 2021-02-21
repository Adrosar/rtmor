package core

// Options ...
type Options struct {
	ShowLogs        bool
	ProxyServerAddr string
	ConfigFile      string
}

// NewOptions ...
func NewOptions() *Options {
	return &Options{}
}
