package configs


type Config struct {
	DSN           string `toml:"dsn"`
}

func NewConfig() *Config {
	return &Config{}
}