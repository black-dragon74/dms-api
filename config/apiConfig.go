package config

import "fmt"

// APIConfig reflects the type for file `config.toml`
type APIConfig struct {
	redis            bool
	monitorDataStore bool
	host             string
	port             int
}

func (a APIConfig) UseRedis() bool {
	return a.redis
}

func (a APIConfig) MonitorDataStore() bool {
	return a.monitorDataStore
}

func (a APIConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
