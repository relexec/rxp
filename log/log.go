package log

import (
	"log"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"

	"github.com/relexec/rxp/config"
)

const (
	LevelDebug = 4
	LevelInfo  = 1
	LevelError = 0
)

// FromConfig returns the [logr.Logger] configured for the rxp library.
func FromConfig(c config.Config) logr.Logger {
	opts := stdr.Options{}
	stdr.SetVerbosity(levelStrToVerbosity(c.Logging.Level))
	l := log.Default()
	return stdr.NewWithOptions(l, opts).WithName("rxp")
}

func levelStrToVerbosity(level string) int {
	switch strings.ToLower(level) {
	case "info":
		return LevelInfo
	case "debug":
		return LevelDebug
	default:
		return LevelError
	}
}
