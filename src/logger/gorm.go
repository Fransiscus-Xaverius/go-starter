package logger

import (
	"context"
	"strings"
	"time"

	appContext "github.com/cde/go-example/src/context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type gormSingleLineLogger struct {
	logger.Interface
}

func (l gormSingleLineLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	// Collapse all newlines/tabs/spaces
	cleanSQL := strings.Join(strings.Fields(sql), " ")

	contextBuilder := appContext.NewContextBuilder(ctx)
	entry := contextBuilder.GetLogger()
	//entry := generic_gorm.GetLoggerFromContext(ctx)

	entry.Printf("[%.2fms] [rows:%d] %s", float64(elapsed.Microseconds())/1000.0, rows, cleanSQL)
}

func MakeGormSingleLineLogger(logger logger.Interface) logger.Interface {
	return gormSingleLineLogger{logger}
}

func MakeGormSingleLineLoggerWithDefaultInterface() logger.Interface {
	baseLogger := logger.New(
		log.New(),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	return gormSingleLineLogger{Interface: baseLogger}
}
