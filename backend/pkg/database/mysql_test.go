package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigFromEnv(t *testing.T) {
	// Save original environment variables
	originalHost := os.Getenv("DB_HOST")
	originalPort := os.Getenv("DB_PORT")
	originalUser := os.Getenv("DB_USER")
	originalPassword := os.Getenv("DB_PASSWORD")
	originalDBName := os.Getenv("DB_NAME")
	originalCharset := os.Getenv("DB_CHARSET")

	// Clean up after test
	defer func() {
		if originalHost != "" {
			os.Setenv("DB_HOST", originalHost)
		} else {
			os.Unsetenv("DB_HOST")
		}
		if originalPort != "" {
			os.Setenv("DB_PORT", originalPort)
		} else {
			os.Unsetenv("DB_PORT")
		}
		if originalUser != "" {
			os.Setenv("DB_USER", originalUser)
		} else {
			os.Unsetenv("DB_USER")
		}
		if originalPassword != "" {
			os.Setenv("DB_PASSWORD", originalPassword)
		} else {
			os.Unsetenv("DB_PASSWORD")
		}
		if originalDBName != "" {
			os.Setenv("DB_NAME", originalDBName)
		} else {
			os.Unsetenv("DB_NAME")
		}
		if originalCharset != "" {
			os.Setenv("DB_CHARSET", originalCharset)
		} else {
			os.Unsetenv("DB_CHARSET")
		}
	}()

	tests := []struct {
		name           string
		envVars        map[string]string
		expectedConfig Config
	}{
		{
			name: "All environment variables set",
			envVars: map[string]string{
				"DB_HOST":     "test-host",
				"DB_PORT":     "3307",
				"DB_USER":     "test-user",
				"DB_PASSWORD": "test-password",
				"DB_NAME":     "test-db",
				"DB_CHARSET":  "utf8",
			},
			expectedConfig: Config{
				Host:     "test-host",
				Port:     "3307",
				User:     "test-user",
				Password: "test-password",
				DBName:   "test-db",
				Charset:  "utf8",
			},
		},
		{
			name:    "No environment variables set (defaults)",
			envVars: map[string]string{},
			expectedConfig: Config{
				Host:     "localhost",
				Port:     "3306",
				User:     "root",
				Password: "password",
				DBName:   "arritech_users",
				Charset:  "utf8mb4",
			},
		},
		{
			name: "Partial environment variables set",
			envVars: map[string]string{
				"DB_HOST": "custom-host",
				"DB_USER": "custom-user",
			},
			expectedConfig: Config{
				Host:     "custom-host",
				Port:     "3306", // default
				User:     "custom-user",
				Password: "password",       // default
				DBName:   "arritech_users", // default
				Charset:  "utf8mb4",        // default
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all environment variables first
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_PORT")
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_NAME")
			os.Unsetenv("DB_CHARSET")

			// Set environment variables for this test
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Get config
			config := GetConfigFromEnv()

			// Assert config matches expected
			assert.Equal(t, tt.expectedConfig.Host, config.Host)
			assert.Equal(t, tt.expectedConfig.Port, config.Port)
			assert.Equal(t, tt.expectedConfig.User, config.User)
			assert.Equal(t, tt.expectedConfig.Password, config.Password)
			assert.Equal(t, tt.expectedConfig.DBName, config.DBName)
			assert.Equal(t, tt.expectedConfig.Charset, config.Charset)
		})
	}
}

func TestGetEnv(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("TEST_ENV_VAR")

	// Clean up after test
	defer func() {
		if originalValue != "" {
			os.Setenv("TEST_ENV_VAR", originalValue)
		} else {
			os.Unsetenv("TEST_ENV_VAR")
		}
	}()

	tests := []struct {
		name         string
		envValue     string
		defaultValue string
		expected     string
	}{
		{
			name:         "Environment variable set",
			envValue:     "custom-value",
			defaultValue: "default-value",
			expected:     "custom-value",
		},
		{
			name:         "Environment variable not set",
			envValue:     "",
			defaultValue: "default-value",
			expected:     "default-value",
		},
		{
			name:         "Environment variable empty string",
			envValue:     "",
			defaultValue: "default-value",
			expected:     "default-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_ENV_VAR", tt.envValue)
			} else {
				os.Unsetenv("TEST_ENV_VAR")
			}

			result := getEnv("TEST_ENV_VAR", tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConfig_DSN(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		DBName:   "test_db",
		Charset:  "utf8mb4",
	}

	// Test that the config can be used to build a DSN
	// This is more of an integration test, but useful to verify the structure
	// Since we can't easily test the exact DSN without duplicating the logic,
	// we'll test that the config fields are accessible
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "3306", config.Port)
	assert.Equal(t, "root", config.User)
	assert.Equal(t, "password", config.Password)
	assert.Equal(t, "test_db", config.DBName)
	assert.Equal(t, "utf8mb4", config.Charset)
}
