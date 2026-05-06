package config

import (
	"github.com/spf13/pflag"
)

// WithOption modifies a Config returned from New.
type WithOption func(*Config)

// New returns a new Config.
func New(opts ...WithOption) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithFlags binds the supplied flags to the Config's fields.
func WithFlags(fs *pflag.FlagSet) WithOption {
	return func(c *Config) {
		c.BindFlags(fs)
	}
}

// Config contains configuration options for the Temporal Cluster Operator
type Config struct {
	// Logging contains options for configuring logging.
	Logging LoggingConfig `json:"logging"`
}

// Validate checks for invalid settings.
func (c *Config) Validate() error {
	return c.Logging.Validate()
}

// BindFlags bings the supplied flagset to the Config's fields.
func (c *Config) BindFlags(fs *pflag.FlagSet) {
	c.Logging.BindFlags(fs)
}
