package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config with all the envs and other possible configurations/parameters
type Config struct {
	Timeout            time.Duration `envconfig:"TIMEOUT" default:"1m"`
	HttPort            int           `envconfig:"HTTP_PORT" default:"8080"`
	HttpReadTimeout    time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"10s"`
	HttpWriteTimeout   time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"10s"`
	HttpMaxHeaderBytes int           `envconfig:"HTTP_MAX_HEADER_BYTES" default:"1048576"`
}

// FromEnvironment returns specification loaded either from an `.env` file or as environ
func FromEnvironment() (Config, error) {
	var c Config
	err := envconfig.Process("booking-insights", &c)
	return c, err
}
