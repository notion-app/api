package log

import (
	"github.com/Sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

// Init initializes the logging functionality of notion
func Init() {
	logger = logrus.New()
}

// Info logs a string with no accompanying metadata
func Info(s string) {
	logger.Info(s)
}

// InfoFields logs a string and a set of fields
func InfoFields(s string, fields map[string]interface{}) {
	logger.WithFields(fields).Info(s)
}

// Error checks if an error passed in is not null, and if it exists prints it
func Error(e error) bool {
	if e != nil {
		logger.Error(e.Error())
		return true
	}
	return false
}
