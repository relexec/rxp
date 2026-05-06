package config

import (
	"github.com/spf13/pflag"
)

const (
	DefaultLoggingLevel  = "info"
	flagLoggingLevel     = "logging-level"
	flagLoggingLevelDesc = "The log level. The default is info. The options are, in descending order of chattiness: debug, info, error"
)

// LoggingConfig contains logging configuration options for rxp.
type LoggingConfig struct {
	// Level is the log level to use.
	//
	// The default is "info". The options are, in descending order of
	// chattiness: "debug", "info", "error".
	Level string `json:"level"`
}

// Validate checks for invalid settings.
func (c *LoggingConfig) Validate() error {
	return nil
}

// BindFlags binds the supplied flagset to the LoggingConfig's fields.
func (c *LoggingConfig) BindFlags(fs *pflag.FlagSet) {
	pflag.StringVar(
		&c.Level,
		flagLoggingLevel,
		DefaultLoggingLevel,
		flagLoggingLevelDesc,
	)
}
