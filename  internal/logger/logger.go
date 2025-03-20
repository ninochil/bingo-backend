package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

func Info(message string) {
	log.Info(message)
}

func Error(err error) {
	log.Error(err)
}