package log

import (
	"github.com/Sirupsen/logrus"
)

var (
	logger = logrus.New()
)

// Init initializes the logging functionality of notion
func Init() {

}

// Info logs a string with no accompanying metadata
func Info(s string) {
	logger.Info(s)
}

// InfoFields logs a string and a set of fields
func InfoFields(s string, fields map[string]interface{}) {
	logger.WithFields(fields).Info(s)
}
