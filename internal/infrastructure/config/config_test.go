package config_test

import (
	"booking-insights/internal/infrastructure/config"
	"os"
	"testing"
	"time"
)

func Test_LoadEnvs(t *testing.T) {

	tests := []struct {
		name                   string
		envs                   map[string]string
		wantHttPort            int
		wantHttpReadTimeout    time.Duration
		wantHttpWriteTimeout   time.Duration
		wantHttpMaxHeaderBytes int
	}{
		{
			name:                   "default envs when not declared",
			wantHttPort:            8080,
			wantHttpReadTimeout:    10 * time.Second,
			wantHttpWriteTimeout:   10 * time.Second,
			wantHttpMaxHeaderBytes: 1048576,
		},
		{
			name: "should match expected envs declared",
			envs: map[string]string{
				"HTTP_PORT":             "3000",
				"HTTP_READ_TIMEOUT":     "20s",
				"HTTP_WRITE_TIMEOUT":    "15s",
				"HTTP_MAX_HEADER_BYTES": "123476",
			},
			wantHttPort:            3000,
			wantHttpReadTimeout:    20 * time.Second,
			wantHttpWriteTimeout:   15 * time.Second,
			wantHttpMaxHeaderBytes: 123476,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()

			if tt.envs != nil {
				for k, v := range tt.envs {
					os.Setenv(k, v)
				}
			}

			cfg, err := config.FromEnvironment()
			if err != nil {
				t.Fatalf("got unexpected error: %s ", err.Error())
			}

			if want, got := tt.wantHttPort, cfg.HttPort; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}
			if want, got := tt.wantHttpReadTimeout, cfg.HttpReadTimeout; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}

			if want, got := tt.wantHttpWriteTimeout, cfg.HttpWriteTimeout; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}

			if want, got := tt.wantHttpMaxHeaderBytes, cfg.HttpMaxHeaderBytes; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}
		})
	}
}
