package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// DefaultFormatter is a custom logrus Formatter
type DefaultFormatter struct{}

// Format ..
func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}

// SetOutput declares the output format
func SetOutput(output string) {
	switch {
	case output == "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		logrus.WithFields(logrus.Fields{}).Info("logger using json config")
	case output == "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		logrus.SetOutput(os.Stdout)
		logrus.WithFields(logrus.Fields{}).Info("logger using text config")
	default:
		logrus.SetFormatter(new(DefaultFormatter))
	}

	levelENV, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		levelENV = "info"
	}
	logLevel, err := logrus.ParseLevel(levelENV)
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
}
