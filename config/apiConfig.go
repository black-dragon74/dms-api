package config

// APIConfig reflects the type for file `config.toml`
type APIConfig struct {
	useRedis bool
}

func (a APIConfig) UseRedis() bool {
	return a.useRedis
}
