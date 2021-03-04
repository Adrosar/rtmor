package core

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/fatih/color"
)

// LogMaster ...
type LogMaster struct {
	ShowAll bool
	Logger  *log.Logger
	Map     map[byte]bool
}

// NewLogMaster ...
func NewLogMaster() *LogMaster {
	var w io.Writer

	if runtime.GOOS == "windows" {
		w = color.Output
	} else {
		w = os.Stdout
	}

	lm := &LogMaster{
		Logger: log.New(w, "", log.Ltime|log.Ldate),
		Map:    make(map[byte]bool),
	}

	return lm
}

// Parse ...
func (lm *LogMaster) Parse(flags string) {
	if flags == "all" {
		lm.ShowAll = true
		return
	}

	for i := range flags {
		lm.Map[flags[i]] = true
	}
}

// Print ...
func (lm *LogMaster) Print(flag byte, v ...interface{}) {
	show := lm.Map[flag]
	if show || lm.ShowAll {
		lm.Logger.Print(v...)
	}
}
