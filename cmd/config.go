package main

import "github.com/cristalhq/aconfig"

type ApplicationConfig struct {
	ListenAddress string `default:":3000" env:"LISTEN_ADDRESS"`
	DatabaseURL   string `env:"DATABASE_URL" required:"true"`
}

func LoadConfig() (*ApplicationConfig, error) {
	config := &ApplicationConfig{}

	// Load configuration from environment variables
	loader := aconfig.LoaderFor(config, aconfig.Config{})
	if err := loader.Load(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *ApplicationConfig) GetListenAddress() string {
	return c.ListenAddress
}
