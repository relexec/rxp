package config

import (
	"os"

	"github.com/spf13/pflag"
)

const (
	DefaultLogLevel  = "info"
	flagLogLevel     = "rxp-logging-level"
	flagLogLevelDesc = "The log level. The default is info. The options are, in descending order of chattiness: debug, info, error"
	EnvVarLogLevel   = "RXP_LOG_LEVEL"
)

// LogConfig contains logging configuration options for rxp.
type LogConfig struct {
	// Level is the log level to use.
	//
	// The default is "info". The options are, in descending order of
	// chattiness: "debug", "info", "error".
	Level string `json:"level"`
}

func (c *LogConfig) SetDefaults() {
	if c.Level == "" {
		lvl := os.Getenv(EnvVarLogLevel)
		if lvl == "" {
			lvl = DefaultLogLevel
		}
		c.Level = lvl
	}
}

// Validate checks for invalid settings.
func (c *LogConfig) Validate() error {
	return nil
}

// BindFlags binds the supplied flagset to the LogConfig's fields.
func (c *LogConfig) BindFlags(fs *pflag.FlagSet) {
	pflag.StringVar(
		&c.Level,
		flagLogLevel,
		DefaultLogLevel,
		flagLogLevelDesc,
	)
}
