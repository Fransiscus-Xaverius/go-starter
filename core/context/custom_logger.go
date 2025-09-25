package context

import (
	"os"
	"time"

	"github.com/cde/go-example/config"
	log "github.com/sirupsen/logrus"
)

var (
	customLogger = &log.Logger{
		Out: os.Stdout,
		//Formatter: &log.JSONFormatter{
		//	TimestampFormat: time.RFC3339Nano,
		//	PrettyPrint:     false,
		//},
		Formatter: &log.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		},
		Hooks:        make(log.LevelHooks),
		Level:        log.ErrorLevel,
		ExitFunc:     os.Exit,
		ReportCaller: true,
	}
)

func NewLogger() *log.Logger {
	if config.Get().AppDebug {
		customLogger.SetLevel(log.InfoLevel)
	}
	return customLogger
}
