package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/samber/do"
)

// LogLevel represents the severity level of a log message as a string
type LogLevel string

// Define log levels as constants
const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

// GetLogLevelValue returns the numeric value of the log level for comparison
func GetLogLevelValue(level LogLevel) int {
	switch level {
	case DEBUG:
		return 0
	case INFO:
		return 1
	case WARN:
		return 2
	case ERROR:
		return 3
	case FATAL:
		return 4
	default:
		return 1 // Default to INFO level
	}
}

// Logger defines the interface for logging operations
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	SetLevel(level LogLevel)
	GetLevel() LogLevel
}

// StandardLogger implements the Logger interface
type StandardLogger struct {
	level  LogLevel
	output io.Writer
}

// NewLogger creates a new StandardLogger with the log level from environment or default
func NewLogger(i *do.Injector) (Logger, error) {
	return &StandardLogger{
		level:  GetLogLevelFromEnv(),
		output: os.Stdout,
	}, nil
}

// NewLoggerWithLevel creates a new StandardLogger with the specified log level
func NewLoggerWithLevel(level LogLevel) Logger {
	return &StandardLogger{
		level:  level,
		output: os.Stdout,
	}
}

// NewLoggerWithOutput creates a new StandardLogger with custom output writer
func NewLoggerWithOutput(level LogLevel, output io.Writer) Logger {
	return &StandardLogger{
		level:  level,
		output: output,
	}
}

// GetLogLevelFromEnv retrieves the log level from environment variable or returns default
func GetLogLevelFromEnv() LogLevel {
	levelStr := os.Getenv("TRANSACTION_LOG_LEVEL")
	if levelStr == "" {
		return INFO // Default level if not specified
	}

	levelStr = strings.ToUpper(levelStr)
	switch LogLevel(levelStr) {
	case DEBUG, INFO, WARN, ERROR, FATAL:
		return LogLevel(levelStr)
	default:
		return INFO
	}
}

// SetLevel changes the current log level
func (l *StandardLogger) SetLevel(level LogLevel) {
	l.level = level
}

// GetLevel returns the current log level
func (l *StandardLogger) GetLevel() LogLevel {
	return l.level
}

// log handles the actual logging with the specified level
func (l *StandardLogger) log(level LogLevel, format string, args ...interface{}) {
	// Skip if the current level is higher than the message level
	if GetLogLevelValue(level) < GetLogLevelValue(l.level) {
		return
	}

	// Get current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// Format message
	message := format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}

	// Create log entry
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)

	// Write to output
	fmt.Fprint(l.output, logEntry)

	// Exit program if it's a FATAL message
	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a debug message
func (l *StandardLogger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs an info message
func (l *StandardLogger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a warning message
func (l *StandardLogger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs an error message
func (l *StandardLogger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal logs a fatal message and exits the program
func (l *StandardLogger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}

// ParseLogLevel converts a string to a LogLevel
func ParseLogLevel(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO // Default to INFO level
	}
}
