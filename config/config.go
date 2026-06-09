package config

import (
	"os"

	"github.com/spf13/pflag"
)

const (
	flagSystemUUID     = "rxp-system-uuid"
	flagSystemUUIDDesc = "Contains the rxp host system UUID. If empty, the value of RXP_SYSTEM_UUID environs variable is used."
	envVarSystemUUID   = "RXP_SYSTEM_UUID"
	flagSystemTag      = "rxp-system-tag"
	flagSystemTagDesc  = "Contains the rxp host system tag. If empty, the value of RXP_SYSTEM_TAG environs variable is used."
	envVarSystemTag    = "RXP_SYSTEM_TAG"
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

// Config contains common configuration options for rxp.
type Config struct {
	// SystemUUID contains the rxp host system UUID.
	SystemUUID string `json:"system_uuid,omitempty"`
	// SystemTag contains the rxp host system Name, if any.
	SystemTag string `json:"system_name,omitempty"`
	// Log contains options for configuring logging.
	Log LogConfig `json:"logging"`
}

// SetDefaults sets any missing values to their defaults or environs variable
// values.
func (c *Config) SetDefaults() {
	if c.SystemUUID == "" {
		c.SystemUUID = os.Getenv(envVarSystemUUID)
	}
	if c.SystemTag == "" {
		c.SystemTag = os.Getenv(envVarSystemTag)
	}
	c.Log.SetDefaults()
}

// Validate checks for invalid settings.
func (c *Config) Validate() error {
	return c.Log.Validate()
}

// BindFlags bings the supplied flagset to the Config's fields.
func (c *Config) BindFlags(fs *pflag.FlagSet) {
	pflag.StringVar(
		&c.SystemUUID,
		flagSystemUUID,
		"",
		flagSystemUUIDDesc,
	)
	pflag.StringVar(
		&c.SystemTag,
		flagSystemTag,
		"",
		flagSystemTagDesc,
	)
	c.Log.BindFlags(fs)
}
