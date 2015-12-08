package log

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

type (
	Fields logrus.Fields
)

var (
	logger *logrus.Logger
)

// Init initializes the logging functionality of notion
func Init() {
	logger = logrus.New()
	// logger.Formatter = &NotionFormatter{}
	// mgoHook, err := NewMongoHook(config.LoggingMongoURL(), "logs")
	// if err != nil {
	// 	fmt.Printf("%v\n", err.Error())
	// 	return
	// }
	// logger.Hooks.Add(mgoHook)
}

// Info logs a string with no accompanying metadata
func Info(s string, args ...interface{}) {
	if len(args) == 0 {
		logger.Info(s)
	} else {
		logger.Info(fmt.Sprintf(s, args...))
	}
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
