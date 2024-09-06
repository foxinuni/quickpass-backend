package core

import "github.com/cristalhq/aconfig"

type ApplicationConfig struct {
	ListenAddress    string `env:"LISTEN_ADDRESS" default:":3000"`
	MigrationsSource string `env:"MIGRATIONS_SRC" default:"file://migrations"`
	DatabaseURL      string `env:"DATABASE_URL" required:"true"`
	JwtSecret        string `env:"JWT_SECRET" required:"true"`
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

func (c *ApplicationConfig) GetMigrationsSource() string {
	return c.MigrationsSource
}

func (c *ApplicationConfig) GetDatabaseURL() string {
	return c.DatabaseURL
}

func (c *ApplicationConfig) GetJwtSecret() string {
	return c.JwtSecret
}
