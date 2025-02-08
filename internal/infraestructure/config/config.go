package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config with all the envs and other possible configurations/parameters
type Config struct {
	HttPort            int `env:"HTTP_PORT" envDefault:"8080"`
	HttpReadTimeout    int `env:"HTTP_READ_TIMEOUT_SEC" envDefault:"10"`
	HttpWriteTimeout   int `env:"HTTP_WRITE_TIMEOUT_SEC" envDefault:"10"`
	HttpMaxHeaderBytes int `env:"HTTP_MAX_HEADER_BYTES" envDefault:"1048576"`
}

// LoadEnvParseConfig loads and expports the envs from the .env` file,
// After that, it parses the rnbd inyo the config struct.
func LoadEnvParseConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing ennv variables: %w", err)
	}

	return cfg, nil
}
