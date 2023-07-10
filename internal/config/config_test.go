package config_test

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jtprogru/owlops/internal/config"
)

func TestGetConfig(t *testing.T) {
	// Prepare test environment
	originalArgs := os.Args
	originalEnv := os.Getenv(config.EnvConfigPathName)
	defer func() {
		os.Args = originalArgs
		os.Setenv(config.EnvConfigPathName, originalEnv)
	}()

	// Set up test cases
	testCases := []struct {
		name          string
		args          []string
		env           map[string]string
		expectedError bool
	}{
		{
			name:          "No config file specified",
			args:          []string{"command"},
			env:           map[string]string{},
			expectedError: true,
		},
		{
			name:          "Config file specified as flag",
			args:          []string{"command", "-config=config/config.local.yaml"},
			env:           map[string]string{},
			expectedError: false,
		},
		{
			name:          "Config file specified as environment variable",
			args:          []string{"command"},
			env:           map[string]string{config.EnvConfigPathName: "config/config.local.yaml"},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set command-line arguments
			os.Args = tc.args

			// Set environment variables
			for key, value := range tc.env {
				os.Setenv(key, value)
			}

			// Reset flag
			flag.CommandLine = flag.NewFlagSet(tc.args[0], flag.ExitOnError)

			// Test GetConfig
			conf := config.GetConfig()

			// Assert expectations
			if tc.expectedError {
				assert.Nil(t, conf)
			} else {
				assert.NotNil(t, conf)
				assert.Equal(t, "DEBUG", conf.LogLevel)
				assert.Equal(t, uint16(8081), conf.Port)
			}
		})
	}
}
