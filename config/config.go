package config

// Config holds the configuration for the DIGIT client
type Config struct {
	APIEndpoint string
	AuthToken   string
}

// NewConfig creates a new configuration instance
func NewConfig(apiEndpoint, authToken string) *Config {
	return &Config{
		APIEndpoint: apiEndpoint,
		AuthToken:   authToken,
	}
} 