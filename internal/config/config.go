package config

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogLevel int `toml:"log_level"`
	} `toml:"general"`

	PostgreSQL struct {
		DSN         string `toml:"dsn"`
		Automigrate bool   `toml:"automigrate"`
	} `toml:"postgresql"`

	ExternalAPI struct {
		Bind      string `toml:"bind"`
		TLSCert   string `toml:"tls_cert"`
		TLSKey    string `toml:"tls_key"`
		JWTSecret string `toml:"jwt_secret"`
	} `toml:"external_api"`
}

// C holds the global configuration.
var C Config
