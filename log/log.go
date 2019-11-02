package log

import (
	"github.com/sirupsen/logrus"
)

func CreateLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}
