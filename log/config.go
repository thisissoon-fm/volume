package log

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	vLevel              = "log.level"
	vFormat             = "log.format"
	vFile               = "log.logfile"
	vConsoleOutput      = "log.console_output"
	vLogstashFormatType = "logstash.type"
)

// Set logging configuration defaults
func init() {
	// Defaults
	viper.SetDefault(vLevel, "info")
	viper.SetDefault(vFormat, "text")
	viper.SetDefault(vFile, "")
	viper.SetDefault(vConsoleOutput, true)
	viper.SetDefault(vLogstashFormatType, "volumesvc")
	// Bind Environment Variables
	viper.BindEnv(
		vLevel,
		vFormat,
		vFile,
		vConsoleOutput,
		vLogstashFormatType)
}

// Bind CLI Flag for log level override
func BindLogLevelFlag(flag *pflag.Flag) {
	viper.BindPFlag(vLevel, flag)
}

// Configuration interface
type Configurer interface {
	Level() string
	File() string
	ConsoleOutput() bool
	Format() string
	LogstashFormatType() string
}

// Default Config Type
type Config struct{}

// Returns the log level from viper
func (c Config) Level() string {
	return viper.GetString(vLevel)
}

// Rrturns a path to a log file to write log messages too
func (c Config) File() string {
	return viper.GetString(vFile)
}

// Returns boolean indicating if we should log to console or not
func (c Config) ConsoleOutput() bool {
	return viper.GetBool(vConsoleOutput)
}

// Returns log format
func (c Config) Format() string {
	return viper.GetString(vFormat)
}

// Returns logstash type
func (c Config) LogstashFormatType() string {
	return viper.GetString(vLogstashFormatType)
}
