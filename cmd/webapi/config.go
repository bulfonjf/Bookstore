package main

const (
	// DefaultDSN is the default datasource name.
	DefaultDSN = "./infrastructure/sqlite/db"
)

type Config struct {
	DB struct {
		DSN string
	}

	HTTP struct {
		Addr     string
		Domain   string
		HashKey  string
		BlockKey string
	}

	GitHub struct {
		ClientID     string
		ClientSecret string
	}
}

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	config.DB.DSN = DefaultDSN
	return config
}
