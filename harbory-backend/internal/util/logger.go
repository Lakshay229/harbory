package util

import (
	"log"
)

type Logger interface {
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
}

type SimpleLogger struct{}

func NewLogger() Logger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) Info(msg string, keyvals ...interface{}) {
	log.Printf("INFO: %s %v\n", msg, keyvals)
}

func (l *SimpleLogger) Error(msg string, keyvals ...interface{}) {
	log.Printf("ERROR: %s %v\n", msg, keyvals)
}

func (l *SimpleLogger) Debug(msg string, keyvals ...interface{}) {
	log.Printf("DEBUG: %s %v\n", msg, keyvals)
}
