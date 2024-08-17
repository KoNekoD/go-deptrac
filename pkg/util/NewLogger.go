package util

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

func NewLogger(module string) *log.Logger {
	opts := log.Options{
		ReportTimestamp: true,
		Prefix:          module,
		TimeFormat:      time.DateTime,
		Level:           log.DebugLevel,
	}
	logger := log.NewWithOptions(os.Stdout, opts)
	return logger
}
