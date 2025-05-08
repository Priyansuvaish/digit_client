package configdigit

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

type Gobalconfig struct {
	defaultconfInfo *Config
}

var globalConfig *Gobalconfig // Singleton instance

// Correct singleton accessor (package-level function)
func GetGlobalConfig() *Gobalconfig {
	if globalConfig == nil {
		globalConfig = &Gobalconfig{}
	}
	return globalConfig
}

// Updated initialization method
func (gc *Gobalconfig) Initialize(apiEndpoint, authToken string) {
	gc.defaultconfInfo = NewConfig(apiEndpoint, authToken)
}

func (dc *Gobalconfig) GetInfo() *Config {
	return dc.defaultconfInfo
}
