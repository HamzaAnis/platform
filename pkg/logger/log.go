package logger

import (
	"os"

	"github.com/op/go-logging"
)

func (l *loggerImpl) Log() *logging.Logger {
	return l.logger
}

func configure(appName string) (*logging.Logger, error) {
	var logger logging.LeveledBackend
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.MustStringFormatter(
		`%{longfile} %{color:bold} ▶ [%{level:.6s}] %{message}%{color:reset}`,
	)

	backendFormatter := logging.NewBackendFormatter(backend, formatter)
	logger = logging.SetBackend(backendFormatter)

	logger.SetLevel(logging.DEBUG, "")

	return logging.GetLogger(appName)
}
