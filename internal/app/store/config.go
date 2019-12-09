package store


// Config ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
	/*LogLevel string `toml:"log_level"`
	Store *store.Config*/
}

func NewConfig() *Config{
	return  &Config{
		/*BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),*/
	}
}
