package xlib

import (
	"io"
	"os"

	"github.com/op/go-logging"
)

var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	// `%{color}[%{level:.4s}]%{color:reset} %{message}`,
)

// GetLogger returns a Logger of go-logging package
func GetLogger(programName string) *logging.Logger {
	log := logging.MustGetLogger(programName)
	return log
}

func init() {
	var stderr io.Writer = os.Stderr
	backend := logging.NewLogBackend(stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, logFormat)
	logging.SetBackend(backendFormatter)
}
