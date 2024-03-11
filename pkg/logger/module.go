package logger

import (
	"log"

	"github.com/op/go-logging"
)

type Logger interface {
	Log() *logging.Logger
}

type loggerImpl struct {
	logger *logging.Logger
}

// Assert that *loggerImpl satisfies the Logger interface
var _ Logger = &loggerImpl{}

func NewLogger(appName string) Logger {
	logger, err := configure(appName)
	if err != nil {
		log.Fatal(err)
	}
	return &loggerImpl{
		logger: logger,
	}
}
