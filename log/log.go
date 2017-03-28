package log

import (
	"io/ioutil"
	"strings"

	"volume/build"
	"volume/log/formatters"
	"volume/log/hooks"

	log "github.com/sirupsen/logrus"
)

// Global log entry
var entry = log.NewEntry(log.New())

// Intitlaise standard logger
func init() {
	entry = entry.WithFields(log.Fields{
		"version": build.Version(),
	})
}

// Conveneience access to logrus.Fields
type Fields log.Fields

// Initialise
func Initialize() {
	Configure(Config{})
}

// Configure global logger
func Configure(config Configurer) {
	// Set log level
	SetLevel(config.Level())
	// Log to file
	fp := config.File()
	if fp != "" {
		LogToFile(fp)
	}
	// Disable Console Output
	if !config.ConsoleOutput() {
		entry.Logger.Out = ioutil.Discard
	}
	// Set Format
	SetFormat(config.Format(), map[string]interface{}{
		"logstash.type": config.LogstashFormatType(),
	})
}

// Set the Log Level
func SetLevel(lvl string) {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		entry.Logger.Level = log.DebugLevel
	case "info":
		entry.Logger.Level = log.InfoLevel
	case "warn":
		entry.Logger.Level = log.WarnLevel
	case "error":
		entry.Logger.Level = log.ErrorLevel
	default:
		entry.Logger.Level = log.InfoLevel
	}
}

// Set the format of the log
func SetFormat(format string, args map[string]interface{}) {
	format = strings.ToLower(format)
	switch format {
	case "json":
		entry.Logger.Formatter = &log.JSONFormatter{}
	case "logstash":
		typ, ok := args["logstash.type"].(string)
		if !ok {
			typ = "xerosyncsvc"
		}
		entry.Logger.Formatter = &formatters.LogstashFormatter{
			Type: typ,
		}
	default:
		entry.Logger.Formatter = &log.TextFormatter{
			FullTimestamp: true,
		}
	}
}

// Log to a file
func LogToFile(path string) {
	entry.Logger.Hooks.Add(hooks.NewFileHook(path))
}

// Add Error
func WithError(err error) *log.Entry {
	return entry.WithError(err)
}

// Add one field to the log context
func WithField(k string, v interface{}) *log.Entry {
	return entry.WithField(k, v)
}

// Log multiple fields
func WithFields(f Fields) *log.Entry {
	return entry.WithFields(log.Fields(f))
}

// Debug logging
func Debug(format string, args ...interface{}) {
	entry.Debugf(format, args...)
}

// Info Logging
func Info(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

// Warn logging
func Warn(format string, args ...interface{}) {
	entry.Warnf(format, args...)
}

// Error Logging
func Error(format string, args ...interface{}) {
	entry.Errorf(format, args...)
}
