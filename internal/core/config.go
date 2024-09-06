package core

import (
	"github.com/cristalhq/aconfig"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation"
)

type ApplicationConfig struct {
	ListenAddress    string `env:"LISTEN_ADDRESS" default:":3000"`
	MigrationsSource string `env:"MIGRATIONS_SRC" default:"file://migrations"`
	DatabaseURL      string `env:"DATABASE_URL" required:"true"`
	JwtSecret        string `env:"JWT_SECRET" required:"true"`
	SendgridEmail    string `env:"SENDGRID_EMAIL" required:"true"`
	SendgridAPIKey   string `env:"SENDGRID_API_KEY" required:"true"`
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

// --- QuickpassAPIOptions implementation --- //
var _ presentation.QuickpassAPIOptions = (*ApplicationConfig)(nil)

func (c *ApplicationConfig) GetListenAddress() string {
	return c.ListenAddress
}

// --- PgStoreFactoryOptions implementation --- //
var _ PgStoreFactoryOptions = (*ApplicationConfig)(nil)

func (c *ApplicationConfig) GetMigrationsSource() string {
	return c.MigrationsSource
}

func (c *ApplicationConfig) GetDatabaseURL() string {
	return c.DatabaseURL
}

// --- JwtOptions implementation --- //
var _ services.JwtAuthServiceOptions = (*ApplicationConfig)(nil)

func (c *ApplicationConfig) GetJwtSecret() string {
	return c.JwtSecret
}

// --- SendgridOptions implementation --- //
var _ services.SendgridEmailServiceOptions = (*ApplicationConfig)(nil)

func (c *ApplicationConfig) GetSendgridEmail() string {
	return c.SendgridEmail
}

func (c *ApplicationConfig) GetSendgridAPIKey() string {
	return c.SendgridAPIKey
}
