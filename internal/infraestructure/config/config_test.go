package config_test

import (
	"booking-req-insights/internal/infraestructure/config"
	"bytes"
	"os"
	"testing"
)

const dotEnvFileName = ".env"

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(t *testing.T) func(t *testing.T) {
	// testsDir := t.TempDir()

	t.Log("create tmp .env file")

	// envFile, err := os.Create(testsDir + "/" + dotEnvFileName)
	envFile, err := os.Create(dotEnvFileName)
	if err != nil {
		t.Fatalf("got unexpected err: %s", err)
	}

	// Return a function to teardown the test
	return func(tb *testing.T) {
		t.Log("teardown suite")
		if err = envFile.Close(); err != nil {
			t.Fatalf("got unexpected err: %s", err)
		}
		if err = os.Remove(envFile.Name()); err != nil {
			t.Fatalf("got unexpected err: %s", err)
		}
	}
}

func Test_LoadEnvs(t *testing.T) {

	tests := []struct {
		name                   string
		envs                   map[string]string
		wantHttPort            int
		wantHttpReadTimeout    int
		wantHttpWriteTimeout   int
		wantHttpMaxHeaderBytes int
	}{
		{
			name:                   "should get default envs when not declared",
			wantHttPort:            8080,
			wantHttpReadTimeout:    10,
			wantHttpWriteTimeout:   10,
			wantHttpMaxHeaderBytes: 1048576,
		},
		{
			name: "should match expected envs declared",
			envs: map[string]string{
				"HTTP_PORT":              "3000",
				"HTTP_READ_TIMEOUT_SEC":  "20",
				"HTTP_WRITE_TIMEOUT_SEC": "15",
				"HTTP_MAX_HEADER_BYTES":  "123476",
			},
			wantHttPort:            3000,
			wantHttpReadTimeout:    20,
			wantHttpWriteTimeout:   15,
			wantHttpMaxHeaderBytes: 123476,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tearDown := setupSuite(t)
			defer tearDown(t)

			envFile, err := os.OpenFile(dotEnvFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				t.Fatalf("got unexpected err: %s", err)
			}

			// envs := map[string]string{"HTTP_PORT": "3000"}
			if tt.envs != nil {
				var envsStr string
				for k, v := range tt.envs {
					envsStr += k + "=" + v + "\n"

				}

				buff := bytes.NewBufferString(envsStr)
				if _, err := envFile.Write(buff.Bytes()); err != nil {
					t.Fatalf("got unexpected err: %s", err)
				}

			}

			cfg, err := config.LoadEnvParseConfig()
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
