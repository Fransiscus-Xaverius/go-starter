package context

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	customLogger = &log.Logger{
		Out: os.Stdout,
		Formatter: &log.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			PrettyPrint:     false,
		},
		//Formatter:    &log.TextFormatter{},
		Hooks:        make(log.LevelHooks),
		Level:        log.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
)

func NewLogger() *log.Logger {
	return customLogger
}
