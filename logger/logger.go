package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger(env string) *logrus.Logger {
	logger := logrus.New()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.SetOutput(os.Stdout)
		logger.Warn("Failed to log to file, using default stdout")
	} else {
		logger.SetOutput(io.MultiWriter(os.Stdout, file))
	}

	if env == "production" {
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return logger
}
