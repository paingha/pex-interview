package lib

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	TitleField       = "Title"
	MessageField     = "Message"
	TypeField        = "Type"
	OriginField      = "Origin"
	RequestIDField   = "Request ID"
	RequestPathField = "Request Path"
	InfoType         = "info"
	ErrorType        = "error"
	FatalType        = "fatal"
)

// Log
type Log struct {
	*logrus.Logger
	RequestID     string
	RequestPath   string
	RequestMethod string
}

// NewLogger returns a new logger.
func NewLogger() *Log {
	return &Log{
		logrus.New(),
		"",
		"",
		"",
	}
}

// NewRequestID generates a new request id for the log message.
func NewRequestID() string {
	return uuid.New().String()
}

// WithRequestID attaches a request id to be added to the log message.
func (l *Log) WithRequestID() *Log {
	l.RequestID = NewRequestID()
	return l
}

// WithRequestPath attaches a request path to be added to the log message.
func (l *Log) WithRequestPath(requestPath string) *Log {
	l.RequestPath = requestPath
	return l
}

// WithRequestMethod attaches a request method to be added to the log message.
func (l *Log) WithRequestMethod(requestMethod string) *Log {
	l.RequestMethod = requestMethod
	return l
}

// FormatErrorMessage utility function to format an error message.
func FormatErrorMessage(message string, err error) string {
	return fmt.Sprintf("%s - %v", message, err)
}

// FileZone returns the filename that calls the logger. It helps trace issues the logger logs.
func FileZone() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

// Info logs an information message.
func (l *Log) Info(title, message, origin string) {
	l.WithFields(logrus.Fields{
		TitleField:       title,
		MessageField:     message,
		TypeField:        InfoType,
		OriginField:      origin,
		RequestIDField:   l.RequestID,
		RequestPathField: l.RequestPath,
	}).Info(message)
}

// Warn logs a warning message.
func (l *Log) Warn(title, message, origin string) {
	l.WithFields(logrus.Fields{
		TitleField:       title,
		MessageField:     message,
		TypeField:        ErrorType,
		OriginField:      origin,
		RequestIDField:   l.RequestID,
		RequestPathField: l.RequestPath,
	}).Warn(message)
}

// Error logs an error message.
func (l *Log) Error(title, message, origin string) {
	l.WithFields(logrus.Fields{
		TitleField:       title,
		MessageField:     message,
		TypeField:        ErrorType,
		OriginField:      origin,
		RequestIDField:   l.RequestID,
		RequestPathField: l.RequestPath,
	}).Error(message)
}

// Fatal logs an error message and panics.
func (l *Log) Fatal(title, message, origin string) {
	l.WithFields(logrus.Fields{
		TitleField:       title,
		MessageField:     message,
		TypeField:        ErrorType,
		OriginField:      origin,
		RequestIDField:   l.RequestID,
		RequestPathField: l.RequestPath,
	}).Fatal(message)
}
